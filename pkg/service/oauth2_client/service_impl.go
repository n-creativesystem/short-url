package oauth2client

import (
	"context"
	"crypto/subtle"
	"errors"

	"github.com/go-oauth2/oauth2/v4"
	oauth2client "github.com/n-creativesystem/short-url/pkg/domain/oauth2_client"
	"github.com/n-creativesystem/short-url/pkg/domain/repository"
	"github.com/n-creativesystem/short-url/pkg/domain/tx"
	"github.com/n-creativesystem/short-url/pkg/service"
	"github.com/n-creativesystem/short-url/pkg/utils"
)

type serviceImpl struct {
	repo     oauth2client.Repository
	beginner tx.ContextBeginner
}

func newServiceImpl(repo oauth2client.Repository, beginner tx.ContextBeginner) *serviceImpl {
	return &serviceImpl{repo: repo, beginner: beginner}
}

func NewService(repo oauth2client.Repository, beginner tx.ContextBeginner) Service {
	return newServiceImpl(repo, beginner)
}

func (impl *serviceImpl) GetByID(ctx context.Context, clientId string) (oauth2.ClientInfo, error) {
	if v, err := impl.repo.GetByID(ctx, clientId); err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return nil, service.ErrNotFound
		}
		return nil, service.Wrap(err, "Service oauth2_client")
	} else {
		return v, nil
	}
}

func (impl *serviceImpl) FindAll(ctx context.Context, user string) ([]oauth2client.Client, error) {
	if v, err := impl.repo.Find(ctx, user); err == nil {
		return v, nil
	} else {
		return nil, service.Wrap(err, "OAuth2")
	}
}

func (impl *serviceImpl) FindByID(ctx context.Context, id, user string) (*oauth2client.Client, error) {
	if v, err := impl.repo.FindByID(ctx, id, user); err == nil {
		return v, nil
	} else {
		return nil, service.Wrap(err, "OAuth2")
	}
}

func (impl *serviceImpl) RegisterClient(ctx context.Context, user, appName string) (RegisterResult, error) {
	var result RegisterResult
	clientId, err := utils.GenerateRandomString(20)
	if err != nil {
		return result, service.Wrap(err, "Service oauth2_client")
	}
	clientSecret, err := utils.GenerateRandomString(40)
	if err != nil {
		return result, service.Wrap(err, "Service oauth2_client")
	}
	client := oauth2client.NewClient(
		clientId, clientSecret, "", false, user, appName,
	)
	if err := impl.repo.Create(ctx, &client); err != nil {
		return result, service.Wrap(err, "Service oauth2_client")
	}
	result.ClientId = clientId
	result.ClientSecret = clientSecret
	return result, nil
}

func (impl *serviceImpl) UpdateClient(ctx context.Context, id, user, appName string) error {
	return impl.repo.Update(ctx, id, user, appName)
}

func (impl *serviceImpl) DeleteClient(ctx context.Context, user, clientId string) error {
	info, err := impl.repo.GetByID(ctx, clientId)
	if err != nil {
		return service.Wrap(err, "Service oauth2_client")
	}
	if subtle.ConstantTimeCompare([]byte(user), []byte(info.GetUserID())) != 1 {
		return service.Wrap(errors.New("Cannot delete because the credentials are incorrect."), "Service oauth2_client")
	}
	if err := impl.repo.Delete(ctx, user, clientId); err != nil {
		return service.Wrap(err, "Service oauth2_client")
	}
	return nil
}
