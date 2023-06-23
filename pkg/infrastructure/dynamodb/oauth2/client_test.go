package oauth2

import (
	"testing"

	"github.com/n-creativesystem/short-url/pkg/infrastructure/dynamodb/tests"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/shared_test/oauth2"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials/crypto"
	"github.com/stretchr/testify/require"
)

func TestOAuth2Client(t *testing.T) {
	require := require.New(t)
	err := crypto.Init()
	require.NoError(err)
	ctx, rollback, err := tests.SetUpDynamoDB(
		tests.Table{Name: clientTable, KeyColumn: "id"},
		tests.Table{Name: accessTokenTable, KeyColumn: "id"},
		tests.Table{Name: refreshTokenTable, KeyColumn: "id"},
		tests.Table{Name: authorizationTable, KeyColumn: "id"},
	)
	require.NoError(err)
	defer rollback()
	repoImpl := NewOAuth2Client()
	oauth2.TestOAuth2Client(t, ctx, repoImpl)
}

func TestOAuth2ClientError(t *testing.T) {
	require := require.New(t)
	err := crypto.Init()
	require.NoError(err)
	ctx, _, err := tests.SetUpDynamoDB()
	require.NoError(err)
	repoImpl := NewOAuth2Client()
	oauth2.TestOAuth2ClientError(t, ctx, repoImpl)
}
