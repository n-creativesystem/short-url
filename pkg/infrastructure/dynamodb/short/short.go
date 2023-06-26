package short

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/n-creativesystem/short-url/pkg/domain/repository"
	"github.com/n-creativesystem/short-url/pkg/domain/short"
	infra "github.com/n-creativesystem/short-url/pkg/infrastructure/dynamodb"
	"github.com/n-creativesystem/short-url/pkg/utils"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials"
	"golang.org/x/sync/errgroup"
)

type repositoryImpl struct {
}

var (
	_ short.Repository = (*repositoryImpl)(nil)
)

func newRepositoryImpl() *repositoryImpl {
	return &repositoryImpl{}
}

func NewRepository() short.Repository {
	return newRepositoryImpl()
}

func (impl *repositoryImpl) Get(ctx context.Context, key string) (*short.Short, error) {
	svc := infra.GetExecutor(ctx)
	resp, err := svc.GetItem(ctx, &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			shortColumns.Key: &types.AttributeValueMemberS{Value: key},
		},
		TableName: aws.String(tableName),
	})
	if err != nil {
		return nil, err
	}
	if resp.Item == nil {
		return nil, repository.ErrRecordNotFound
	}
	result := toModel(resp.Item)
	return result.Short, nil
}

func (impl *repositoryImpl) Put(ctx context.Context, value short.Short) error {
	svc := infra.GetExecutor(ctx)
	now := utils.NowString()
	_, err := svc.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]types.AttributeValue{
			shortColumns.Key:       &types.AttributeValueMemberS{Value: value.GetKey()},
			shortColumns.URL:       &types.AttributeValueMemberS{Value: value.GetEncryptURL().MustEncrypt()},
			shortColumns.Author:    &types.AttributeValueMemberS{Value: value.GetAuthor()},
			shortColumns.CreatedAt: &types.AttributeValueMemberS{Value: now},
			shortColumns.UpdatedAt: &types.AttributeValueMemberS{Value: now},
		},
	})
	return err
}

func (impl *repositoryImpl) Del(ctx context.Context, key, author string) (bool, error) {
	const maxBatchSize = 25
	svc := infra.GetExecutor(ctx)
	keys, err := impl.existsByKeyAndAuthor(ctx, key, author)
	if err != nil {
		return false, err
	}
	var eg errgroup.Group
	for i := 0; i < len(keys); i += maxBatchSize {
		end := i + maxBatchSize
		if end > len(keys) {
			end = len(keys)
		}
		batchKeys := keys[i:end]
		eg.Go(func() error {
			return deleteBatch(svc, ctx, batchKeys)
		})
	}
	if err := eg.Wait(); err != nil {
		return false, err
	}
	return true, nil
}

func (impl *repositoryImpl) Exists(ctx context.Context, key string) (bool, error) {
	v, err := impl.Get(ctx, key)
	if err != nil {
		var exists *types.ResourceNotFoundException
		if errors.As(err, &exists) {
			return false, nil
		}
		if errors.Is(err, repository.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return v != nil, nil
}

func (impl *repositoryImpl) FindAll(ctx context.Context, author string) ([]short.ShortWithTimeStamp, error) {
	filter := "#col1 = :author"
	input := dynamodb.ScanInput{
		TableName:        aws.String(tableName),
		FilterExpression: &filter,
		ExpressionAttributeNames: map[string]string{
			"#col1": "author",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":author": &types.AttributeValueMemberS{Value: author},
		},
	}
	return impl.findAll(ctx, input)
}

func (impl *repositoryImpl) FindByKeyAndAuthor(ctx context.Context, key, author string) (*short.ShortWithTimeStamp, error) {
	filter := "#col1 = :key and #col2 = :author"
	input := dynamodb.ScanInput{
		TableName:        aws.String(tableName),
		FilterExpression: &filter,
		ExpressionAttributeNames: map[string]string{
			"#col1": "key",
			"#col2": "author",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":key":    &types.AttributeValueMemberS{Value: key},
			":author": &types.AttributeValueMemberS{Value: author},
		},
	}
	shorts, err := impl.findAll(ctx, input)
	if err != nil {
		return nil, err
	}
	return &shorts[0], nil
}

func (impl *repositoryImpl) existsByKeyAndAuthor(ctx context.Context, key, author string) ([]string, error) {
	filter := "#col1 = :key and #col2 = :author"
	input := dynamodb.ScanInput{
		TableName:        aws.String(tableName),
		FilterExpression: &filter,
		ExpressionAttributeNames: map[string]string{
			"#col1": "key",
			"#col2": "author",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":key":    &types.AttributeValueMemberS{Value: key},
			":author": &types.AttributeValueMemberS{Value: author},
		},
	}
	shorts, err := impl.findAll(ctx, input)
	if err != nil {
		var exists *types.ResourceNotFoundException
		if errors.As(err, &exists) {
			return nil, repository.ErrRecordNotFound
		}
		return nil, err
	}
	var keys []string
	for _, short := range shorts {
		keys = append(keys, short.GetKey())
	}
	return keys, nil
}

func (impl *repositoryImpl) findAll(ctx context.Context, input dynamodb.ScanInput) ([]short.ShortWithTimeStamp, error) {
	svc := infra.GetExecutor(ctx)
	var eKey map[string]types.AttributeValue
	values := make([]short.ShortWithTimeStamp, 0, 10)
	for {
		if eKey != nil {
			input.ExclusiveStartKey = eKey
		}
		out, err := svc.Scan(ctx, &input)
		if err != nil {
			return nil, err
		}
		for _, item := range out.Items {
			v := toModel(item)
			values = append(values, *v)
		}
		if out.LastEvaluatedKey != nil {
			eKey = out.LastEvaluatedKey
		} else {
			break
		}
	}
	return values, nil
}

func deleteBatch(svc *dynamodb.Client, ctx context.Context, keys []string) error {
	table := tableName
	var writeRequests []types.WriteRequest
	for _, key := range keys {
		writeRequests = append(writeRequests, types.WriteRequest{
			DeleteRequest: &types.DeleteRequest{
				Key: map[string]types.AttributeValue{
					shortColumns.Key: &types.AttributeValueMemberS{Value: key},
				},
			},
		})
	}
	_, err := svc.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			table: writeRequests,
		},
	})
	return err
}

func toModel(item map[string]types.AttributeValue) *short.ShortWithTimeStamp {
	keyValue, _ := item[shortColumns.Key].(*types.AttributeValueMemberS)
	urlValue, _ := item[shortColumns.URL].(*types.AttributeValueMemberS)
	authorValue, _ := item[shortColumns.Author].(*types.AttributeValueMemberS)
	createdAt, _ := item[shortColumns.CreatedAt].(*types.AttributeValueMemberS)
	updatedAt, _ := item[shortColumns.CreatedAt].(*types.AttributeValueMemberS)
	value := short.NewShort(credentials.NewEncryptStringWithMustDecrypt(urlValue.Value).UnmaskedString(), keyValue.Value, authorValue.Value)
	return &short.ShortWithTimeStamp{
		Short:     value,
		CreatedAt: utils.StringToTime(createdAt.Value),
		UpdatedAt: utils.StringToTime(updatedAt.Value),
	}
}
