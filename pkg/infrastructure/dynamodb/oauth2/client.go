package oauth2

import (
	"context"
	"crypto/subtle"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/go-oauth2/oauth2/v4"
	oauth2client "github.com/n-creativesystem/short-url/pkg/domain/oauth2_client"
	"github.com/n-creativesystem/short-url/pkg/domain/repository"
	infra "github.com/n-creativesystem/short-url/pkg/infrastructure/dynamodb"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials/crypto"
	"github.com/n-creativesystem/short-url/pkg/utils/logging"
)

type oauth2ClientImpl struct{}

var (
	_ oauth2client.Repository = (*oauth2ClientImpl)(nil)
)

func NewOAuth2Client() oauth2client.Repository {
	return newOAuth2Client()
}

func newOAuth2Client() *oauth2ClientImpl {
	return &oauth2ClientImpl{}
}

func (impl *oauth2ClientImpl) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	if id == "" {
		return nil, repository.ErrRecordNotFound
	}
	return impl.getItem(ctx, clientTable, id)
}

func (impl *oauth2ClientImpl) Create(ctx context.Context, client *oauth2client.Client) error {
	id := client.GetID()
	buf, err := fromClient(client)
	if err != nil {
		return err
	}
	svc := infra.GetExecutor(ctx)
	_, err = svc.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(clientTable),
		Item: map[string]types.AttributeValue{
			clientColumns.Id:   &types.AttributeValueMemberS{Value: id},
			clientColumns.User: &types.AttributeValueMemberS{Value: client.GetUserID()},
			clientColumns.Data: &types.AttributeValueMemberB{Value: buf},
		},
	})
	return err
}

func (impl *oauth2ClientImpl) Delete(ctx context.Context, user, id string) error {
	info, err := impl.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if subtle.ConstantTimeCompare([]byte(user), []byte(info.GetUserID())) != 1 {
		return repository.ErrRecordNotFound
	}
	svc := infra.GetExecutor(ctx)
	_, err = svc.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(clientTable),
		Key: map[string]types.AttributeValue{
			clientColumns.Id: &types.AttributeValueMemberS{Value: id},
		},
	})
	return err
}

func (impl *oauth2ClientImpl) getItem(ctx context.Context, tableName string, id string) (*oauth2client.Client, error) {
	svc := infra.GetExecutor(ctx)
	out, err := svc.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &tableName,
		Key: map[string]types.AttributeValue{
			clientColumns.Id: &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return nil, err
	}
	if out == nil || out.Item == nil {
		return nil, repository.ErrRecordNotFound
	}
	buf, ok := out.Item[clientColumns.Data].(*types.AttributeValueMemberB)
	if !ok {
		return nil, repository.ErrRecordNotFound
	}
	return toClient(buf.Value)
}

func (impl *oauth2ClientImpl) Find(ctx context.Context, user string) ([]oauth2client.Client, error) {
	results := make([]oauth2client.Client, 0, 100)
	svc := infra.GetExecutor(ctx)
	filterEx := expression.Name("user").Equal(expression.Value(user))
	expr, err := expression.NewBuilder().WithFilter(filterEx).Build()
	if err != nil {
		return nil, err
	}
	paginator := dynamodb.NewScanPaginator(svc, &dynamodb.ScanInput{
		TableName:                 aws.String(clientTable),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, item := range output.Items {
			v, ok := item[clientColumns.Data].(*types.AttributeValueMemberB)
			if !ok {
				continue
			}
			client, err := toClient(v.Value)
			if err != nil {
				logging.Default().Warn(err)
				continue
			}
			results = append(results, *client)
		}
	}
	return results, nil
}

func (impl *oauth2ClientImpl) FindByID(ctx context.Context, id, user string) (*oauth2client.Client, error) {
	if id == "" {
		return nil, repository.ErrRecordNotFound
	}
	v, err := impl.getItem(ctx, clientTable, id)
	if err != nil {
		return nil, err
	}
	if v.GetUserID() != user {
		return nil, repository.ErrRecordNotFound
	}
	return v, nil
}

func (impl *oauth2ClientImpl) Update(ctx context.Context, id, user, appName string) error {
	value, err := impl.FindByID(ctx, id, user)
	if err != nil {
		return err
	}
	value.AppName = appName
	buf, err := fromClient(value)
	if err != nil {
		return err
	}
	update := expression.Set(expression.Name(clientColumns.Data), expression.Value(buf))
	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		return err
	}
	svc := infra.GetExecutor(ctx)
	_, err = svc.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(clientTable),
		Key: map[string]types.AttributeValue{
			clientColumns.Id: &types.AttributeValueMemberS{Value: id},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})
	return err
}

func fromClient(client *oauth2client.Client) ([]byte, error) {
	buf, err := json.Marshal(client)
	if err != nil {
		return nil, err
	}
	encrypted, err := crypto.Encrypt(string(buf))
	if err != nil {
		return nil, err
	}
	return []byte(encrypted), nil
}

func toClient(buf []byte) (*oauth2client.Client, error) {
	if len(buf) == 0 {
		return nil, repository.ErrRecordNotFound
	}
	decrypted, err := crypto.Decrypt(string(buf))
	if err != nil {
		return nil, err
	}
	var info oauth2client.Client
	if err := json.Unmarshal([]byte(decrypted), &info); err != nil {
		return nil, err
	}
	return &info, nil
}
