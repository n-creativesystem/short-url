package oauth2

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/dynamodb/tests"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/shared_test/oauth2"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials/crypto"
	"github.com/stretchr/testify/require"
)

func TestOAuth2Token(t *testing.T) {
	require := require.New(t)
	err := crypto.Init()
	require.NoError(err)
	ctx, rollback, err := tests.SetUpDynamoDB(
		tests.Table{Name: accessTokenTable, KeyColumn: "id"},
		tests.Table{Name: authorizationTable, KeyColumn: "id"},
		tests.Table{Name: refreshTokenTable, KeyColumn: "id"},
	)
	require.NoError(err)
	defer rollback()
	tokenStore := NewOAuth2Token(1, nil)
	oauth2.TestOAuth2Token(t, ctx, tokenStore)
}
