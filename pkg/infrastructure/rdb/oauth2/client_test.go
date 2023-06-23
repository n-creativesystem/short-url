package oauth2

import (
	"context"
	"testing"

	"github.com/n-creativesystem/short-url/pkg/domain/config"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/tests"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/shared_test/oauth2"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials/crypto"
	"github.com/stretchr/testify/require"
)

func TestOAuth2Client(t *testing.T) {
	require := require.New(t)
	err := crypto.Init()
	require.NoError(err)
	t.Parallel()
	for _, d := range config.RDB {
		d := d
		config.SetDriver(d)
		entClient := tests.GetTestDBAndMigrate("shorturl_by_oauth2_client")
		rdb.SetClient(entClient)
		ctx := context.Background()
		repoImpl := NewOAuthClient()
		oauth2.TestOAuth2Client(t, ctx, repoImpl)
	}
}

func TestOAuth2ClientError(t *testing.T) {
	require := require.New(t)
	err := crypto.Init()
	require.NoError(err)
	for _, d := range config.RDB {
		d := d
		config.SetDriver(d)
		client := tests.GetTestDBAndMigrate("shorturl_by_oauth2_client")
		rdb.SetClient(client)
		ctx := context.Background()
		repoImpl := NewOAuthClient()
		oauth2.TestOAuth2ClientError(t, ctx, repoImpl)
	}
}
