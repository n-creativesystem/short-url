package short

import (
	"context"
	"testing"

	"github.com/n-creativesystem/short-url/pkg/infrastructure/aws"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/dynamodb"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/dynamodb/migration"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/dynamodb/tests"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/shared_test/short"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials/crypto"
	"github.com/stretchr/testify/require"
)

func testSetDynamoDB() (context.Context, error) {
	ctx := context.Background()
	endpoint := aws.NewEndpoint()
	awsConfig, err := aws.NewConfig(ctx, endpoint.EndpointResolver())
	if err != nil {
		return nil, err
	}
	svc, err := dynamodb.NewDynamoDB(ctx, dynamodb.WithAwsConfig(awsConfig))
	if err != nil {
		return nil, err
	}
	dynamodb.SetDB(svc)
	if err := migration.Up(ctx, svc, migration.WithIgnore()); err != nil {
		return nil, err
	}
	return ctx, nil
}

func TestShort(t *testing.T) {
	require := require.New(t)
	err := crypto.Init()
	require.NoError(err)
	ctx, rollback, err := tests.SetUpDynamoDB(tests.Table{Name: tableName, KeyColumn: shortColumns.ID})
	require.NoError(err)
	defer rollback()
	repoImpl := NewRepository()
	short.TestShort(t, ctx, repoImpl)
}
