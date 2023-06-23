package migration

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/dynamodb/tables"
)

func shortCreateInput() *dynamodb.CreateTableInput {
	return &dynamodb.CreateTableInput{
		TableName: aws.String(tables.Short),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("key"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("key"),
				KeyType:       types.KeyTypeHash,
			},
		},
		BillingMode: types.BillingModePayPerRequest,
	}
}

func shortDeleteInput() *dynamodb.DeleteTableInput {
	return &dynamodb.DeleteTableInput{
		TableName: aws.String(tables.Short),
	}
}
