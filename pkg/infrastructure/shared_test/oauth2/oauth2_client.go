package oauth2

import (
	"context"
	"testing"

	oauth2client "github.com/n-creativesystem/short-url/pkg/domain/oauth2_client"
	"github.com/n-creativesystem/short-url/pkg/domain/repository"
	"github.com/stretchr/testify/require"
)

func TestOAuth2Client(t *testing.T, ctx context.Context, repoImpl oauth2client.Repository) {
	require := require.New(t)
	test := struct {
		clientId     string
		clientSecret string
		domain       string
		public       bool
		userId       string
		appName      string
	}{
		"client_id", "client_secret", "http://localhost", false, "test", "app_name",
	}
	modelClient := oauth2client.NewClient(
		test.clientId, test.clientSecret, test.domain, test.public, test.userId, test.appName,
	)
	err := repoImpl.Create(ctx, &modelClient)
	require.NoError(err)

	info, err := repoImpl.GetByID(ctx, modelClient.GetID())
	require.NoError(err)

	require.Equal(test.clientId, info.GetID())
	require.Equal(test.clientSecret, info.GetSecret())
	require.Equal(test.domain, info.GetDomain())
	require.Equal(test.public, info.IsPublic())
	require.Equal(test.userId, info.GetUserID())

	client, err := repoImpl.FindByID(ctx, modelClient.GetID(), modelClient.GetUserID())
	require.NoError(err)
	require.Equal(test.clientId, info.GetID())
	require.Equal(test.clientSecret, info.GetSecret())
	require.Equal(test.domain, info.GetDomain())
	require.Equal(test.public, info.IsPublic())
	require.Equal(test.userId, info.GetUserID())
	require.Equal(test.appName, client.GetAppName())

	clients, err := repoImpl.Find(ctx, client.GetUserID())
	require.NoError(err)
	require.Len(clients, 1)
	client = &clients[0]
	require.Equal(test.clientId, info.GetID())
	require.Equal(test.clientSecret, info.GetSecret())
	require.Equal(test.domain, info.GetDomain())
	require.Equal(test.public, info.IsPublic())
	require.Equal(test.userId, info.GetUserID())
	require.Equal(test.appName, client.GetAppName())

	err = repoImpl.Delete(ctx, modelClient.GetUserID(), modelClient.GetID())
	require.NoError(err)
}

func TestOAuth2ClientError(t *testing.T, ctx context.Context, repoImpl oauth2client.Repository) {
	require := require.New(t)
	info, err := repoImpl.GetByID(ctx, "")
	require.Error(err)
	require.ErrorIs(err, repository.ErrRecordNotFound)
	require.Nil(info)

	info, err = repoImpl.GetByID(ctx, "no_data")
	require.Error(err)
	require.ErrorIs(err, repository.ErrRecordNotFound)
	require.Nil(info)

	err = repoImpl.Delete(ctx, "no_user", "no_data")
	require.Error(err)
	require.ErrorIs(err, repository.ErrRecordNotFound)

	info, err = repoImpl.GetByID(ctx, "no_data")
	require.Error(err)
	require.Equal("not found", err.Error())
	require.Nil(info)

	err = repoImpl.Delete(ctx, "no_user", "no_data")
	require.Error(err)
	require.Equal("not found", err.Error())

	err = repoImpl.Delete(ctx, "no_user", "no_data")
	require.Error(err)
	require.Equal("not found", err.Error())
}
