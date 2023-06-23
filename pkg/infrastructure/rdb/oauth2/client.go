package oauth2

import (
	"context"

	"github.com/go-oauth2/oauth2/v4"
	domain_oauth2client "github.com/n-creativesystem/short-url/pkg/domain/oauth2_client"
	"github.com/n-creativesystem/short-url/pkg/domain/repository"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent"
	oauth2client "github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent/oauth2client"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent/predicate"
)

func NewOAuthClient() domain_oauth2client.Repository {
	return &oauth2ClientImpl{}
}

type oauth2ClientImpl struct{}

func (s *oauth2ClientImpl) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	return s.findOne(ctx, id)
}

func (s *oauth2ClientImpl) Create(ctx context.Context, client *domain_oauth2client.Client) error {
	db := rdb.GetExecutor(ctx)
	entity := db.OAuth2Client.Create()
	entity.SetID(client.GetID())
	entity.SetSecret(client.GetEncryptSecret())
	entity.SetDomain(client.GetEncryptDomain())
	entity.SetPublic(client.IsPublic())
	entity.SetUserID(client.GetUserID())
	entity.SetAppName(client.GetAppName())
	_, err := entity.Save(ctx)
	return err
}

func (s *oauth2ClientImpl) Delete(ctx context.Context, user, id string) error {
	db := rdb.GetExecutor(ctx)
	result, err := db.OAuth2Client.Delete().Where(
		oauth2client.And(oauth2client.IDEQ(id), oauth2client.UserIDEQ(user)),
	).Exec(ctx)
	if err != nil {
		if rdb.IsNotFoundRecord(err) {
			return repository.ErrRecordNotFound
		}
		return err
	}
	if result == 0 {
		return repository.ErrRecordNotFound
	}
	return nil
}

func (s *oauth2ClientImpl) Find(ctx context.Context, user string) ([]domain_oauth2client.Client, error) {
	db := rdb.GetExecutor(ctx)
	value, err := db.OAuth2Client.Query().Where(oauth2client.UserIDEQ(user)).All(ctx)
	if err != nil {
		return nil, err
	}
	results := make([]domain_oauth2client.Client, len(value))
	for idx, v := range value {
		m := domain_oauth2client.NewClient(
			v.ID,
			v.Secret.UnmaskedString(),
			v.Domain.UnmaskedString(),
			v.Public,
			v.UserID,
			v.AppName,
		)
		results[idx] = m
	}
	return results, nil
}

func (s *oauth2ClientImpl) FindByID(ctx context.Context, id, user string) (*domain_oauth2client.Client, error) {
	if id == "" {
		return nil, repository.ErrRecordNotFound
	}
	return s.findOne(ctx, id, oauth2client.UserIDEQ(user))
}

func (s *oauth2ClientImpl) Update(ctx context.Context, id, user, appName string) error {
	if id == "" {
		return repository.ErrRecordNotFound
	}
	db := rdb.GetExecutor(ctx)
	value := db.OAuth2Client.UpdateOne(&ent.OAuth2Client{ID: id, UserID: user})
	value.SetAppName(appName)
	_, err := value.Save(ctx)
	return err
}

func (s *oauth2ClientImpl) findOne(ctx context.Context, id string, ps ...predicate.OAuth2Client) (*domain_oauth2client.Client, error) {
	if id == "" {
		return nil, repository.ErrRecordNotFound
	}

	db := rdb.GetExecutor(ctx)
	q := db.OAuth2Client.Query().Where(oauth2client.IDEQ(id))
	if len(ps) > 0 {
		q.Where(ps...)
	}
	result, err := q.First(ctx)
	if err != nil {
		if rdb.IsNotFoundRecord(err) {
			return nil, repository.ErrRecordNotFound
		}
		return nil, err
	}
	m := domain_oauth2client.NewClient(
		result.ID,
		result.Secret.UnmaskedString(),
		result.Domain.UnmaskedString(),
		result.Public,
		result.UserID,
		result.AppName,
	)
	return &m, nil
}
