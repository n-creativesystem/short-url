package oauth2

import (
	"context"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/n-creativesystem/short-url/pkg/domain/config"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/tests"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/shared_test/oauth2"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials/crypto"
	"github.com/stretchr/testify/require"
)

func TestOAuth2Token(t *testing.T) {
	require := require.New(t)
	err := crypto.Init()
	require.NoError(err)
	t.Parallel()
	for _, d := range config.RDB {
		d := d
		t.Run(d.String(), func(t *testing.T) {
			config.SetDriver(d)
			client := tests.GetTestDBAndMigrate("shorturl_by_oauth2_token")
			rdb.SetClient(client)
			ctx := context.Background()
			tokenStore := NewOAuth2TokenWithOption(context.Background(), WithGCTimeInterval(1))
			oauth2.TestOAuth2Token(t, ctx, tokenStore)
		})
	}
}
