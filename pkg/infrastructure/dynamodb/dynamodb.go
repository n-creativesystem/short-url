package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewDynamoDB(ctx context.Context, opts ...DynamoDBOption) (*dynamodb.Client, error) {
	cfg := &dynamoDBConfig{}
	for _, opt := range opts {
		opt.apply(cfg)
	}
	svc := dynamodb.NewFromConfig(cfg.cfg)
	return svc, nil
}
