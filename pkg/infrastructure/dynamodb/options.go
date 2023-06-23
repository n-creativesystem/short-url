package dynamodb

import "github.com/aws/aws-sdk-go-v2/aws"

type dynamoDBConfig struct {
	cfg aws.Config
}

type DynamoDBOption interface {
	apply(*dynamoDBConfig)
}

type dynamoDBOptionFn func(*dynamoDBConfig)

func (f dynamoDBOptionFn) apply(opt *dynamoDBConfig) {
	f(opt)
}

func WithAwsConfig(cfg aws.Config) DynamoDBOption {
	return dynamoDBOptionFn(func(dd *dynamoDBConfig) {
		dd.cfg = cfg
	})
}
