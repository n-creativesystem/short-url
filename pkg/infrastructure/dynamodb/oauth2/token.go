package oauth2

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/friendsofgo/errors"
	"github.com/go-oauth2/oauth2/v4"
	oauth2client "github.com/n-creativesystem/short-url/pkg/domain/oauth2_client"
	oauth2token "github.com/n-creativesystem/short-url/pkg/domain/oauth2_token"
	"github.com/n-creativesystem/short-url/pkg/domain/repository"
	infra "github.com/n-creativesystem/short-url/pkg/infrastructure/dynamodb"
	"github.com/n-creativesystem/short-url/pkg/utils/logging"
)

type oauth2TokenImpl struct {
	ticker *time.Ticker
	client oauth2client.Repository
}

var (
	_ oauth2token.Repository = (*oauth2TokenImpl)(nil)
)

func NewOAuth2Token(gcInterval int, repo oauth2client.Repository) oauth2token.Repository {
	return newOAuth2Token(gcInterval, repo)
}

func newOAuth2Token(gcInterval int, repo oauth2client.Repository) *oauth2TokenImpl {
	impl := &oauth2TokenImpl{
		ticker: time.NewTicker(time.Second * time.Duration(gcInterval)),
		client: repo,
	}
	go impl.gc()
	return impl
}

func (impl *oauth2TokenImpl) gc() {
	for range impl.ticker.C {
		ctx := context.Background()
		var wg sync.WaitGroup
		for _, table := range []string{authorizationTable, accessTokenTable, refreshTokenTable} {
			wg.Add(1)
			go func(table string) {
				defer wg.Done()
				ids, err := impl.getDataByExpiration(ctx, table)
				if err != nil {
					impl.error(err)
					return
				}
				if err := impl.deleteByIds(ctx, table, ids); err != nil {
					impl.error(err)
				}
			}(table)
		}
		wg.Wait()
	}
}

func (impl *oauth2TokenImpl) Close() {
	if impl.ticker != nil {
		impl.ticker.Stop()
	}
}

func (impl *oauth2TokenImpl) Create(ctx context.Context, info oauth2.TokenInfo) error {
	input := make(map[string][]types.WriteRequest)
	if impl.client != nil {
		clientInfo, err := impl.client.GetByID(ctx, info.GetClientID())
		if err != nil {
			return err
		}
		info.SetUserID(clientInfo.GetUserID())
	}
	token := oauth2token.NewToken(info)
	data := []byte(token.Encode())
	if code := info.GetCode(); code != "" {
		input[authorizationTable] = []types.WriteRequest{
			{
				PutRequest: &types.PutRequest{
					Item: impl.generatePutItemForAuthorization(info, data),
				},
			},
		}
	} else {
		if code := info.GetAccess(); code != "" {
			input[accessTokenTable] = []types.WriteRequest{
				{
					PutRequest: &types.PutRequest{
						Item: impl.generatePutItemForAccessToken(info, data),
					},
				},
			}
		}
		if refresh := info.GetRefresh(); refresh != "" {
			input[refreshTokenTable] = []types.WriteRequest{
				{
					PutRequest: &types.PutRequest{
						Item: impl.generatePutItemForRefreshToken(info, data),
					},
				},
			}
		}
	}
	svc := infra.GetExecutor(ctx)
	_, err := svc.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
		RequestItems: input,
	})
	return err
}

func (impl *oauth2TokenImpl) RemoveByCode(ctx context.Context, code string) error {
	return impl.deleteItem(ctx, authorizationTable, code)
}

func (impl *oauth2TokenImpl) RemoveByAccess(ctx context.Context, access string) error {
	return impl.deleteItem(ctx, accessTokenTable, access)
}

func (impl *oauth2TokenImpl) RemoveByRefresh(ctx context.Context, refresh string) error {
	return impl.deleteItem(ctx, refreshTokenTable, refresh)
}

func (impl *oauth2TokenImpl) GetByCode(ctx context.Context, code string) (oauth2.TokenInfo, error) {
	return impl.getItem(ctx, authorizationTable, code)
}

func (impl *oauth2TokenImpl) GetByAccess(ctx context.Context, access string) (oauth2.TokenInfo, error) {
	return impl.getItem(ctx, accessTokenTable, access)
}

func (impl *oauth2TokenImpl) GetByRefresh(ctx context.Context, refresh string) (oauth2.TokenInfo, error) {
	return impl.getItem(ctx, refreshTokenTable, refresh)
}

func (impl *oauth2TokenImpl) getItem(ctx context.Context, tableName string, code string) (oauth2.TokenInfo, error) {
	svc := infra.GetExecutor(ctx)
	out, err := svc.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &tableName,
		Key: map[string]types.AttributeValue{
			tokenColumns.Id: &types.AttributeValueMemberS{Value: code},
		},
	})
	if err != nil {
		return nil, err
	}
	if out == nil || out.Item == nil {
		return nil, repository.ErrRecordNotFound
	}
	buf, ok := out.Item[tokenColumns.Data].(*types.AttributeValueMemberB)
	if !ok {
		return nil, repository.ErrRecordNotFound
	}
	tokenInfo, err := oauth2token.Decode(string(buf.Value))
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}

func (impl *oauth2TokenImpl) deleteItem(ctx context.Context, tableName string, code string) error {
	svc := infra.GetExecutor(ctx)
	out, err := svc.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &tableName,
		Key: map[string]types.AttributeValue{
			tokenColumns.Id: &types.AttributeValueMemberS{Value: code},
		},
		ReturnValues: types.ReturnValueAllOld,
	})
	if err != nil {
		return err
	}
	if out == nil || out.Attributes == nil {
		return repository.ErrRecordNotFound
	}
	return nil
}

func (impl *oauth2TokenImpl) generatePutItemForAuthorization(info oauth2.TokenInfo, data []byte) map[string]types.AttributeValue {
	code := info.GetCode()
	expiredAt := info.GetCodeCreateAt().Add(info.GetCodeExpiresIn())
	return generateItem(code, expiredAt, data)
}

func (impl *oauth2TokenImpl) generatePutItemForAccessToken(info oauth2.TokenInfo, data []byte) map[string]types.AttributeValue {
	token := info.GetAccess()
	expiredAt := info.GetAccessCreateAt().Add(info.GetAccessExpiresIn())
	return generateItem(token, expiredAt, data)
}

func (impl *oauth2TokenImpl) generatePutItemForRefreshToken(info oauth2.TokenInfo, data []byte) map[string]types.AttributeValue {
	token := info.GetRefresh()
	expiredAt := info.GetRefreshCreateAt().Add(info.GetRefreshExpiresIn())
	return generateItem(token, expiredAt, data)
}

func (impl *oauth2TokenImpl) getDataByExpiration(ctx context.Context, table string) ([]string, error) {
	svc := infra.GetDB()
	filter := "#col <= :p"
	var eKey map[string]types.AttributeValue
	keys := make([]string, 0, 10)
	for {
		input := dynamodb.ScanInput{
			TableName:        &table,
			FilterExpression: &filter,
			ExpressionAttributeNames: map[string]string{
				"#col": tokenColumns.ExpiredAt,
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":p": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", time.Now().Unix())},
			},
		}
		if eKey != nil {
			input.ExclusiveStartKey = eKey
		}
		out, err := svc.Scan(ctx, &input)
		if err != nil {
			return nil, err
		}
		for _, record := range out.Items {
			k, _ := record[tokenColumns.Id].(*types.AttributeValueMemberS)
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

func (impl *oauth2TokenImpl) deleteByIds(ctx context.Context, table string, ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	svc := infra.GetDB()
	var writeRequests []types.WriteRequest
	for _, id := range ids {
		writeRequests = append(writeRequests, types.WriteRequest{
			DeleteRequest: &types.DeleteRequest{
				Key: map[string]types.AttributeValue{
					tokenColumns.Id: &types.AttributeValueMemberS{Value: id},
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

func (impl *oauth2TokenImpl) error(err error) {
	if err != nil {
		logging.Default().Error(errors.Wrap(err, "[OAUTH2-TOKEN]"))
	}
}

func generateItem(id string, expiredAt time.Time, data []byte) map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		tokenColumns.Id:        &types.AttributeValueMemberS{Value: id},
		tokenColumns.ExpiredAt: &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", expiredAt.Unix())},
		tokenColumns.Data:      &types.AttributeValueMemberB{Value: data},
	}
}
