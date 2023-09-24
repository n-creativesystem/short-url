package user

import (
	"context"
	"errors"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/google/uuid"
	"github.com/n-creativesystem/short-url/pkg/domain/repository"
	"github.com/n-creativesystem/short-url/pkg/domain/social"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent/users"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials"
	"github.com/n-creativesystem/short-url/pkg/utils/hash"
)

type impl struct {
}

func NewRepository() social.UserRepository {
	return newRepository()
}

func newRepository() *impl {
	return &impl{}
}

func (impl *impl) Register(ctx context.Context, user *social.User) (*social.User, error) {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	db := rdb.GetExecutor(ctx)
	var saver interface {
		Save(ctx context.Context) (*ent.Users, error)
	}

	if v, err := impl.findOne(ctx, user.Email); err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			entity := db.Users.Create()
			entity.SetID(uuid.New())
			entity.SetSubject(user.Subject)
			entity.SetProfile(user.Profile)
			entity.SetEmail(credentials.NewEncryptString(user.Email))
			entity.SetEmailHash(hash.NewHash(user.Email))
			entity.SetEmailVerified(user.EmailVerified)
			entity.SetUsername(credentials.NewEncryptString(user.Username))
			entity.SetPicture(user.Picture)
			entity.SetClaims(credentials.NewEncryptString(string(user.GetClaims())))
			saver = entity
		} else {
			return nil, err
		}
	} else {
		entity := db.Users.UpdateOne(v)
		if !user.EqualSubject(v.Subject) {
			entity.SetSubject(user.Subject)
		}
		if !user.EqualEmailVerified(v.EmailVerified) {
			entity.SetEmailVerified(user.EmailVerified)
		}
		if !user.EqualPicture(v.Picture) {
			entity.SetPicture(user.Picture)
		}
		if !user.EqualProfile(v.Profile) {
			entity.SetProfile(user.Profile)
		}
		if !user.EqualUsername(v.Username.UnmaskedString()) {
			entity.SetUsername(credentials.NewEncryptString(user.Username))
		}
		entity.SetClaims(credentials.NewEncryptString(string(user.GetClaims())))
		saver = entity
	}
	if v, err := saver.Save(ctx); err != nil {
		return nil, err
	} else {
		return toModel(v), nil
	}
}

func (impl *impl) Login(ctx context.Context, email string) (*social.User, error) {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	v, err := impl.findOne(ctx, email)
	if err != nil {
		return nil, err
	}
	return toModel(v), nil
}

func (impl *impl) findOne(ctx context.Context, email string) (*ent.Users, error) {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	db := rdb.GetExecutor(ctx)
	v, err := db.Users.Query().Where(users.EmailHashEQ(hash.NewHash(email))).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, repository.ErrRecordNotFound
		}
		return nil, err
	}
	return v, nil
}

func toModel(value *ent.Users) *social.User {
	info := &oidc.UserInfo{
		Subject:       value.Subject,
		Profile:       value.Profile,
		Email:         value.Email.UnmaskedString(),
		EmailVerified: value.EmailVerified,
	}
	u := &social.User{
		UserInfo: info,
		Username: value.Username.UnmaskedString(),
		Picture:  value.Picture,
	}
	u.SetClaims([]byte(value.Claims.UnmaskedString()))
	return u
}
