package migration

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type upInputFunc func() *dynamodb.CreateTableInput
type downInputFunc func() *dynamodb.DeleteTableInput

func Up(ctx context.Context, client *dynamodb.Client, opts ...Option) error {
	opt := NewMigration(opts...)
	inputs := []upInputFunc{
		shortCreateInput,
		oauth2clientCreateInput,
		oauth2authorizationCreateInput,
		oauth2accessTokenCreateInput,
		oauth2refreshTokenCreateInput,
	}
	for _, fn := range inputs {
		input := fn()
		_, err := client.CreateTable(ctx, input)
		if err != nil {
			if opt.Ignore() {
				var exists *types.ResourceInUseException
				if errors.As(err, &exists) {
					continue
				}
			}
			return err
		}
	}
	return nil
}

func Down(ctx context.Context, client *dynamodb.Client, opts ...Option) error {
	opt := NewMigration(opts...)
	inputs := []downInputFunc{
		shortDeleteInput,
	}
	for _, fn := range inputs {
		input := fn()
		_, err := client.DeleteTable(ctx, input)
		if err != nil {
			if opt.Ignore() {
				var exists *types.ResourceNotFoundException
				if errors.As(err, &exists) {
					continue
				}
			}
			return err
		}
	}
	return nil
}
