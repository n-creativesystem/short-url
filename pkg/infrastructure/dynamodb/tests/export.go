package tests

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/aws"
	infra "github.com/n-creativesystem/short-url/pkg/infrastructure/dynamodb"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/dynamodb/migration"
	"github.com/n-creativesystem/short-url/pkg/utils"
	"golang.org/x/sync/errgroup"
)

func SetUpDynamoDB(tables ...Table) (context.Context, func(), error) {
	ctx := context.Background()
	endpoint := aws.NewEndpoint()
	awsConfig, err := aws.NewConfig(ctx, endpoint.EndpointResolver())
	if err != nil {
		return nil, nil, err
	}
	svc, err := infra.NewDynamoDB(ctx, infra.WithAwsConfig(awsConfig))
	if err != nil {
		return nil, nil, err
	}
	infra.SetDB(svc)
	if err := migration.Up(ctx, svc, migration.WithIgnore()); err != nil {
		return nil, nil, err
	}
	tableItems := map[string][]map[string]types.AttributeValue{}
	for _, table := range tables {
		items, err := GetAllData(svc, ctx, table)
		if err != nil {
			return nil, nil, err
		}
		if len(items) > 0 {
			tableItems[table.Name] = items
		}
		_ = TruncateTable(ctx, table)
	}
	rollback := func(tableItems map[string][]map[string]types.AttributeValue) func() {
		return func() {
			for name, items := range tableItems {
				if err := PutBatch(svc, ctx, name, items); err != nil {
					log.Println(err)
				}
			}
		}
	}(tableItems)
	return ctx, rollback, nil
}

type Table struct {
	Name      string
	KeyColumn string
}

func TruncateTable(ctx context.Context, table Table) error {
	const maxBatchSize = 25
	svc := infra.GetDB()
	values, err := GetAllKeys(svc, ctx, table)
	if err != nil {
		var notFound *types.ResourceNotFoundException
		if errors.As(err, &notFound) {
			return nil
		}
		return err
	}
	ids := utils.SplitArray(values, maxBatchSize)
	var eg errgroup.Group
	for _, id := range ids {
		id := id
		eg.Go(func() error {
			return DeleteBatch(svc, ctx, table, id)
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

func GetAllKeys(svc *dynamodb.Client, ctx context.Context, table Table) ([]string, error) {
	var eKey map[string]types.AttributeValue
	keys := make([]string, 0, 10)
	for {
		input := dynamodb.ScanInput{
			TableName: &table.Name,
		}
		if eKey != nil {
			input.ExclusiveStartKey = eKey
		}
		out, err := svc.Scan(ctx, &input)
		if err != nil {
			return nil, err
		}
		for _, record := range out.Items {
			k, _ := record[table.KeyColumn].(*types.AttributeValueMemberS)
			keys = append(keys, k.Value)
		}
		if out.LastEvaluatedKey != nil {
			eKey = out.LastEvaluatedKey
		} else {
			break
		}
	}
	return keys, nil
}

func GetAllData(svc *dynamodb.Client, ctx context.Context, table Table) ([]map[string]types.AttributeValue, error) {
	var eKey map[string]types.AttributeValue
	items := make([]map[string]types.AttributeValue, 0, 10)
	for {
		input := dynamodb.ScanInput{
			TableName: &table.Name,
		}
		if eKey != nil {
			input.ExclusiveStartKey = eKey
		}
		out, err := svc.Scan(ctx, &input)
		if err != nil {
			return nil, err
		}
		items = append(items, out.Items...)
		if out.LastEvaluatedKey != nil {
			eKey = out.LastEvaluatedKey
		} else {
			break
		}
	}
	return items, nil
}

func DeleteBatch(svc *dynamodb.Client, ctx context.Context, table Table, ids []string) error {
	var writeRequests []types.WriteRequest
	for _, id := range ids {
		writeRequests = append(writeRequests, types.WriteRequest{
			DeleteRequest: &types.DeleteRequest{
				Key: map[string]types.AttributeValue{
					table.KeyColumn: &types.AttributeValueMemberS{Value: id},
				},
			},
		})
	}
	_, err := svc.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			table.Name: writeRequests,
		},
	})
	return err
}

func PutBatch(svc *dynamodb.Client, ctx context.Context, table string, items []map[string]types.AttributeValue) error {
	var writeRequests []types.WriteRequest
	for _, item := range items {
		item := item
		writeRequests = append(writeRequests, types.WriteRequest{
			PutRequest: &types.PutRequest{
				Item: item,
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
