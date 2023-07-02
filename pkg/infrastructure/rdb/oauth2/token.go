package oauth2

import (
	"context"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/go-oauth2/oauth2/v4"
	oauth2client "github.com/n-creativesystem/short-url/pkg/domain/oauth2_client"
	domain_oauth2token "github.com/n-creativesystem/short-url/pkg/domain/oauth2_token"
	"github.com/n-creativesystem/short-url/pkg/domain/repository"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent/oauth2token"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent/predicate"
	"github.com/n-creativesystem/short-url/pkg/utils"
	"github.com/n-creativesystem/short-url/pkg/utils/logging"
)

func NewOAuth2Token(gcInterval int, repo oauth2client.Repository) domain_oauth2token.Repository {
	return NewOAuth2TokenWithOption(WithGCTimeInterval(gcInterval), WithOAuth2Client(repo))
}

func NewOAuth2TokenWithOption(opts ...TokenOption) domain_oauth2token.Repository {
	o := &tokenOption{}
	for _, opt := range opts {
		opt.apply(o)
	}
	impl := &tokenImpl{
		tokenOption: o,
	}
	if utils.IsAPI() {
		go impl.gc()
	}
	return impl
}

func isNoRows(err error) bool {
	return rdb.IsNotFoundRecord(err)
}

func checkNoRows(err error) error {
	if err == nil {
		return err
	}
	if isNoRows(err) {
		return nil
	}
	return err
}

type tokenImpl struct {
	*tokenOption
}

func (impl *tokenImpl) Close() {
	impl.ticker.Stop()
}

func (impl *tokenImpl) gc() {
	for range impl.ticker.C {
		impl.clean()
	}
}

func (impl *tokenImpl) clean() {
	impl.error(impl.clear())
}

func (impl *tokenImpl) error(err error) {
	if err != nil {
		logging.Default().Error(errors.Wrap(err, "[OAUTH2-TOKEN]"))
	}
}

func (impl *tokenImpl) Create(ctx context.Context, info oauth2.TokenInfo) error {
	if impl.client != nil {
		clientInfo, err := impl.client.GetByID(ctx, info.GetClientID())
		if err != nil {
			return err
		}
		info.SetUserID(clientInfo.GetUserID())
	}
	token := domain_oauth2token.NewToken(info)
	db := rdb.GetExecutor(ctx)
	item := db.OAuth2Token.Create()
	item.SetData(token.Encode())
	if code := info.GetCode(); code != "" {
		item.SetCode(code)
		item.SetExpiredAt(info.GetCodeCreateAt().Add(info.GetCodeExpiresIn()).Unix())
	} else {
		item.SetAccess(info.GetAccess())
		item.SetExpiredAt(info.GetAccessCreateAt().Add(info.GetAccessExpiresIn()).Unix())
		if refresh := info.GetRefresh(); refresh != "" {
			item.SetRefresh(info.GetRefresh())
			item.SetExpiredAt(info.GetRefreshCreateAt().Add(info.GetRefreshExpiresIn()).Unix())
		}
	}
	_, err := item.Save(ctx)
	return err
}

func (impl *tokenImpl) RemoveByCode(ctx context.Context, code string) error {
	return impl.remove(ctx, oauth2token.CodeEQ(code))
}

func (impl *tokenImpl) RemoveByAccess(ctx context.Context, access string) error {
	return impl.remove(ctx, oauth2token.AccessEQ(access))
}

func (impl *tokenImpl) RemoveByRefresh(ctx context.Context, refresh string) error {
	return impl.remove(ctx, oauth2token.RefreshEQ(refresh))
}

func (impl *tokenImpl) remove(ctx context.Context, ps ...predicate.OAuth2Token) error {
	db := rdb.GetExecutor(ctx)
	_, err := db.OAuth2Token.Delete().Where(ps...).Exec(ctx)
	if checkNoRows(err) != nil {
		return err
	}
	return nil
}

func (impl *tokenImpl) getToken(ctx context.Context, db *ent.Client, ps ...predicate.OAuth2Token) (*ent.OAuth2Token, error) {
	if v, err := db.OAuth2Token.Query().Where(ps...).First(ctx); err != nil {
		if rdb.IsNotFoundRecord(err) {
			return nil, repository.ErrRecordNotFound
		}
		return nil, err
	} else {
		return v, nil
	}
}

func (impl *tokenImpl) dbModelToToken(ctx context.Context, ps ...predicate.OAuth2Token) (oauth2.TokenInfo, error) {
	db := rdb.GetExecutor(ctx)
	token, err := impl.getToken(ctx, db, ps...)
	if err != nil {
		return nil, err
	}
	tokenInfo, err := domain_oauth2token.Decode(token.Data)
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}

func (impl *tokenImpl) GetByCode(ctx context.Context, code string) (oauth2.TokenInfo, error) {
	return impl.dbModelToToken(ctx, oauth2token.CodeEQ(code))
}

func (impl *tokenImpl) GetByAccess(ctx context.Context, access string) (oauth2.TokenInfo, error) {
	return impl.dbModelToToken(ctx, oauth2token.AccessEQ(access))
}

func (impl *tokenImpl) GetByRefresh(ctx context.Context, refresh string) (oauth2.TokenInfo, error) {
	return impl.dbModelToToken(ctx, oauth2token.RefreshEQ(refresh))
}

func (impl *tokenImpl) clear() error {
	ctx := context.Background()
	db := rdb.GetClient()
	now := time.Now().Unix()
	emptyCode := oauth2token.CodeEQ("")
	emptyAccess := oauth2token.AccessEQ("")
	emptyRefresh := oauth2token.RefreshEQ("")
	nowLte := oauth2token.ExpiredAtLTE(now)
	_, err := db.OAuth2Token.Delete().Where(oauth2token.Or(nowLte, oauth2token.And(emptyCode, emptyAccess, emptyRefresh))).Exec(ctx)
	if checkNoRows(err) != nil {
		return err
	}
	return nil
}
