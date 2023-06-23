package oauth2client

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/golang/mock/gomock"
	"github.com/n-creativesystem/short-url/pkg/domain/repository"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/noop"
	mock_oauth2client "github.com/n-creativesystem/short-url/pkg/mock/repository/oauth2client"
	"github.com/n-creativesystem/short-url/pkg/service"
	"github.com/n-creativesystem/short-url/pkg/utils"
	"github.com/n-creativesystem/short-url/pkg/utils/hash"
	"github.com/stretchr/testify/assert"
)

type testTable[T any, R any] struct {
	name          string
	data          T
	prepareMockFn func(repoMock *mock_oauth2client.MockRepository)
	want          func(t *testing.T, result R)
	wantErr       error
}

func testMockRepository(mockCtl *gomock.Controller) *mock_oauth2client.MockRepository {
	repo := mock_oauth2client.NewMockRepository(mockCtl)
	return repo
}

func TestOAuth2ClientGetByID(t *testing.T) {
	ctx := context.Background()
	tests := []testTable[string, oauth2.ClientInfo]{
		{
			name: "success",
			data: "id",
			prepareMockFn: func(repoMock *mock_oauth2client.MockRepository) {
				result := &models.Client{
					ID:     "client_id",
					Secret: "client_secret",
					Domain: "http://localhost",
					Public: false,
					UserID: "test",
				}
				repoMock.EXPECT().GetByID(ctx, "id").Return(result, nil)
			},
			want: func(t *testing.T, result oauth2.ClientInfo) {
				expect := &models.Client{
					ID:     "client_id",
					Secret: "client_secret",
					Domain: "http://localhost",
					Public: false,
					UserID: "test",
				}
				assert.Equal(t, result, expect)
			},
			wantErr: nil,
		},
		{
			name: "not found",
			data: "id",
			prepareMockFn: func(repoMock *mock_oauth2client.MockRepository) {
				repoMock.EXPECT().GetByID(ctx, "id").Return(nil, repository.ErrRecordNotFound)
			},
			want: func(t *testing.T, result oauth2.ClientInfo) {
				assert.Nil(t, result)
			},
			wantErr: service.ErrNotFound,
		},
		{
			name: "other error",
			data: "id",
			prepareMockFn: func(repoMock *mock_oauth2client.MockRepository) {
				repoMock.EXPECT().GetByID(ctx, "id").Return(nil, errors.New("Other error"))
			},
			want: func(t *testing.T, result oauth2.ClientInfo) {
				assert.Nil(t, result)
			},
			wantErr: errors.New("Service oauth2_client: Other error"),
		},
	}
	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()
			mockRepo := testMockRepository(mockCtl)
			tt.prepareMockFn(mockRepo)
			beginner, _ := noop.NewBeginner()
			impl := NewService(mockRepo, beginner)
			result, err := impl.GetByID(ctx, tt.data)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr != nil {
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
			tt.want(t, result)
		})
	}
}

func TestOAuth2ClientRegisterClient(t *testing.T) {
	type registerClient struct {
		id      string
		appName string
	}
	ctx := context.Background()
	tests := []testTable[registerClient, RegisterResult]{
		{
			name: "success",
			data: registerClient{"id", "appName"},
			prepareMockFn: func(repoMock *mock_oauth2client.MockRepository) {
				utils.RandFunc = func(b []byte) (int, error) {
					var value bytes.Buffer
					switch len(b) {
					case 20:
						value.Write([]byte("client_id"))
					case 40:
						value.Write([]byte("client_secret"))
					}
					return value.Read(b)
				}
				repoMock.EXPECT().Create(ctx, gomock.Any()).Return(nil)
			},
			want: func(t *testing.T, result RegisterResult) {
				expect := RegisterResult{
					ClientId:     "636c69656e745f69640000000000000000000000",
					ClientSecret: "636c69656e745f736563726574000000000000000000000000000000000000000000000000000000",
				}
				assert.Equal(t, result, expect)
			},
			wantErr: nil,
		},
		{
			name: "failed for create",
			data: registerClient{"id", "appName"},
			prepareMockFn: func(repoMock *mock_oauth2client.MockRepository) {
				repoMock.EXPECT().Create(ctx, gomock.Any()).Return(errors.New("Error"))
			},
			want: func(t *testing.T, result RegisterResult) {
				expect := RegisterResult{
					ClientId:     "",
					ClientSecret: "",
				}
				assert.Equal(t, result, expect)
			},
			wantErr: errors.New("Service oauth2_client: Error"),
		},
		{
			name: "failed for client_id",
			data: registerClient{"id", "appName"},
			prepareMockFn: func(repoMock *mock_oauth2client.MockRepository) {
				utils.RandFunc = func(b []byte) (int, error) {
					var value bytes.Buffer
					switch len(b) {
					case 20:
						return 0, errors.New("failed for client_id")
					case 40:
						value.Write([]byte("client_secret"))
					}
					return value.Read(b)
				}
			},
			want: func(t *testing.T, result RegisterResult) {
				expect := RegisterResult{
					ClientId:     "",
					ClientSecret: "",
				}
				assert.Equal(t, result, expect)
			},
			wantErr: errors.New("Service oauth2_client: failed for client_id"),
		},
		{
			name: "failed for client_secret",
			data: registerClient{"id", "appName"},
			prepareMockFn: func(repoMock *mock_oauth2client.MockRepository) {
				utils.RandFunc = func(b []byte) (int, error) {
					var value bytes.Buffer
					switch len(b) {
					case 20:
						value.Write([]byte("client_id"))
					case 40:
						return 0, errors.New("failed for client_secret")
					}
					return value.Read(b)
				}
			},
			want: func(t *testing.T, result RegisterResult) {
				expect := RegisterResult{
					ClientId:     "",
					ClientSecret: "",
				}
				assert.Equal(t, result, expect)
			},
			wantErr: errors.New("Service oauth2_client: failed for client_secret"),
		},
	}
	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()
			mockRepo := testMockRepository(mockCtl)
			tt.prepareMockFn(mockRepo)
			beginner, _ := noop.NewBeginner()
			impl := NewService(mockRepo, beginner)
			result, err := impl.RegisterClient(ctx, tt.data.id, tt.data.appName)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr != nil {
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
			tt.want(t, result)
		})
	}
}

func TestOAuth2ClientDeleteClient(t *testing.T) {
	type testData struct {
		user string
		key  string
	}
	user := "anonymous"
	hashUser := hash.Sum([]byte(user))
	ctx := context.Background()
	tests := []testTable[testData, struct{}]{
		{
			name: "success",
			data: testData{
				user: user,
				key:  "client_id",
			},
			prepareMockFn: func(repoMock *mock_oauth2client.MockRepository) {
				result := &models.Client{
					ID:     "client_id",
					Secret: "client_secret",
					Domain: "http://localhost",
					Public: false,
					UserID: hashUser,
				}
				repoMock.EXPECT().GetByID(ctx, "client_id").Return(result, nil)
				repoMock.EXPECT().Delete(ctx, hashUser, "client_id").Return(nil)
			},
			wantErr: nil,
		},
		{
			name: "failed for GetByID",
			data: testData{
				user: "test",
				key:  "client_id",
			},
			prepareMockFn: func(repoMock *mock_oauth2client.MockRepository) {
				repoMock.EXPECT().GetByID(ctx, "client_id").Return(nil, errors.New("Error"))
			},
			wantErr: errors.New("Service oauth2_client: Error"),
		},
		{
			name: "failed for not equal user",
			data: testData{
				user: "test",
				key:  "client_id",
			},
			prepareMockFn: func(repoMock *mock_oauth2client.MockRepository) {
				result := &models.Client{
					ID:     "client_id",
					Secret: "client_secret",
					Domain: "http://localhost",
					Public: false,
					UserID: hashUser,
				}
				repoMock.EXPECT().GetByID(ctx, "client_id").Return(result, nil)
			},
			wantErr: errors.New("Service oauth2_client: Cannot delete because the credentials are incorrect."),
		},
		{
			name: "failed for Delete",
			data: testData{
				user: user,
				key:  "client_id",
			},
			prepareMockFn: func(repoMock *mock_oauth2client.MockRepository) {
				result := &models.Client{
					ID:     "client_id",
					Secret: "client_secret",
					Domain: "http://localhost",
					Public: false,
					UserID: hashUser,
				}
				repoMock.EXPECT().GetByID(ctx, "client_id").Return(result, nil)
				repoMock.EXPECT().Delete(ctx, hashUser, "client_id").Return(errors.New("Other error."))
			},
			wantErr: errors.New("Service oauth2_client: Other error."),
		},
	}
	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()
			mockRepo := testMockRepository(mockCtl)
			tt.prepareMockFn(mockRepo)
			beginner, _ := noop.NewBeginner()
			impl := NewService(mockRepo, beginner)
			err := impl.DeleteClient(ctx, tt.data.user, tt.data.key)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr != nil {
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
		})
	}
}
