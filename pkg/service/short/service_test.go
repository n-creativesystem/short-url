package short

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/n-creativesystem/short-url/fixtures"
	"github.com/n-creativesystem/short-url/pkg/domain/config"
	"github.com/n-creativesystem/short-url/pkg/domain/repository"
	"github.com/n-creativesystem/short-url/pkg/domain/short"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/noop"
	mock_config_repo "github.com/n-creativesystem/short-url/pkg/mock/repository/config"
	mock_short_repo "github.com/n-creativesystem/short-url/pkg/mock/repository/short"
	"github.com/n-creativesystem/short-url/pkg/service"
	"github.com/n-creativesystem/short-url/pkg/tests"
	"github.com/n-creativesystem/short-url/pkg/utils/hash"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testTable[T any, R any] struct {
	name          string
	data          T
	prepareMockFn func(shortRepoMock *mock_short_repo.MockRepository, appConfigMock *mock_config_repo.MockApplicationRepository)
	want          func(t *testing.T, result R)
	wantErr       error
}

func testMockRepository(mockCtl *gomock.Controller) *mock_short_repo.MockRepository {
	repo := mock_short_repo.NewMockRepository(mockCtl)
	return repo
}

func testMockAppRepository(mockCtl *gomock.Controller) *mock_config_repo.MockApplicationRepository {
	return mock_config_repo.NewMockApplicationRepository(mockCtl)
}

func TestServiceForGet(t *testing.T) {
	var (
		mockURL   = "http://localhost:8080/example"
		appConfig = &config.Application{
			RetryGenerateCount: 2,
		}
	)
	ctx := context.Background()
	tests := []testTable[string, string]{
		{
			name: "success",
			data: "key",
			prepareMockFn: func(shortRepoMock *mock_short_repo.MockRepository, appConfigMock *mock_config_repo.MockApplicationRepository) {
				model := short.NewShort(mockURL, "key", "test")
				shortRepoMock.EXPECT().Get(ctx, "key").Return(model, nil)
				appConfigMock.EXPECT().Get(gomock.Any()).Return(appConfig, nil)
			},
			want: func(t *testing.T, result string) {
				assert.Equal(t, mockURL, result)
			},
			wantErr: nil,
		},
		{
			name: "not found",
			data: "not found",
			prepareMockFn: func(shortRepoMock *mock_short_repo.MockRepository, appConfigMock *mock_config_repo.MockApplicationRepository) {
				shortRepoMock.EXPECT().Get(ctx, gomock.Any()).Return(nil, repository.ErrRecordNotFound)
				appConfigMock.EXPECT().Get(gomock.Any()).Return(appConfig, nil)
			},
			want: func(t *testing.T, result string) {
				assert.Empty(t, result)
			},
			wantErr: service.ErrNotFound,
		},
		{
			name: "fail",
			data: "fail",
			prepareMockFn: func(shortRepoMock *mock_short_repo.MockRepository, appConfigMock *mock_config_repo.MockApplicationRepository) {
				shortRepoMock.EXPECT().Get(ctx, gomock.Any()).Return(nil, errors.New("Other error"))
				appConfigMock.EXPECT().Get(gomock.Any()).Return(appConfig, nil)
			},
			want: func(t *testing.T, result string) {
				assert.Empty(t, result)
			},
			wantErr: errors.New("Service shortURL: An error occurred while retrieving the URL.: Other error"),
		},
	}
	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()
			mockRepo := testMockRepository(mockCtl)
			mockAppCfg := testMockAppRepository(mockCtl)
			tt.prepareMockFn(mockRepo, mockAppCfg)
			appCfg, err := mockAppCfg.Get(ctx)
			require.NoError(t, err)
			beginner, _ := noop.NewBeginner()
			impl := NewService(mockRepo, appCfg, beginner)
			url, err := impl.GetURL(ctx, tt.data)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("GetURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr != nil {
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
			tt.want(t, url)
		})
	}
}

func TestServiceForGenerate(t *testing.T) {
	var (
		mockURL   = "http://localhost:8080/example"
		mockKey   = "mock"
		appConfig = &config.Application{
			RetryGenerateCount: 2,
		}
	)
	ctx := context.Background()
	type testData struct {
		key    string
		url    string
		author string
	}
	author := "anonymous"
	hashAuthor := hash.Sum([]byte(author))
	tests := []testTable[testData, string]{
		{
			name: "success",
			data: testData{
				key:    "",
				url:    mockURL,
				author: author,
			},
			prepareMockFn: func(shortRepoMock *mock_short_repo.MockRepository, appConfigMock *mock_config_repo.MockApplicationRepository) {
				short.SetGenerator(func(length ...int) string {
					return mockKey
				})
				shortRepoMock.EXPECT().Exists(ctx, gomock.Any()).Return(false, nil)
				v := short.NewShort(mockURL, "", hashAuthor)
				result := &short.ShortWithTimeStamp{
					Short:     v,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}
				shortRepoMock.EXPECT().Put(ctx, gomock.Eq(*v)).Return(result, nil)
				appConfigMock.EXPECT().Get(gomock.Any()).Return(appConfig, nil)
			},
			want: func(t *testing.T, result string) {
				assert.Equal(t, mockKey, result)
			},
			wantErr: nil,
		},
		{
			name: "success random",
			data: testData{
				key:    "",
				url:    mockURL,
				author: author,
			},
			prepareMockFn: func(shortRepoMock *mock_short_repo.MockRepository, appConfigMock *mock_config_repo.MockApplicationRepository) {
				shortRepoMock.EXPECT().Exists(ctx, gomock.Any()).Return(false, nil)
				v := short.NewShort(mockURL, "", hashAuthor)
				result := &short.ShortWithTimeStamp{
					Short:     v,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}
				shortRepoMock.EXPECT().Put(ctx, tests.NewIgnoreUnexportedFieldsMatcher(*v, "key")).Return(result, nil)
				appConfigMock.EXPECT().Get(gomock.Any()).Return(appConfig, nil)
			},
			want: func(t *testing.T, result string) {
				assert.NotEmpty(t, result)
			},
			wantErr: nil,
		},
		{
			name: "duplicate error",
			data: testData{
				key: "",
				url: mockURL,
			},
			prepareMockFn: func(shortRepoMock *mock_short_repo.MockRepository, appConfigMock *mock_config_repo.MockApplicationRepository) {
				short.SetGenerator(func(length ...int) string {
					return mockKey
				})
				shortRepoMock.EXPECT().Exists(ctx, gomock.Any()).AnyTimes().Return(true, nil)
				appConfigMock.EXPECT().Get(gomock.Any()).Return(appConfig, nil)
			},
			want: func(t *testing.T, result string) {
				assert.Empty(t, result)
			},
			wantErr: errors.New("The number of URL generation attempts reached, but URL could not be generated."),
		},
		{
			name: "validation error by required",
			data: testData{
				key: "",
				url: "",
			},
			prepareMockFn: func(shortRepoMock *mock_short_repo.MockRepository, appConfigMock *mock_config_repo.MockApplicationRepository) {
				short.SetGenerator(func(length ...int) string {
					return ""
				})
				shortRepoMock.EXPECT().Exists(ctx, gomock.Any()).AnyTimes().Return(true, nil)
				appConfigMock.EXPECT().Get(gomock.Any()).Return(appConfig, nil)
			},
			want: func(t *testing.T, result string) {
				assert.Empty(t, result)
			},
			wantErr: errors.New("key: the value is required; url: the value is required."),
		},
		{
			name: "validation error by invalid url",
			data: testData{
				key: "",
				url: "invalid url",
			},
			prepareMockFn: func(shortRepoMock *mock_short_repo.MockRepository, appConfigMock *mock_config_repo.MockApplicationRepository) {
				shortRepoMock.EXPECT().Exists(ctx, gomock.Any()).AnyTimes().Return(true, nil)
				appConfigMock.EXPECT().Get(gomock.Any()).Return(appConfig, nil)
			},
			want: func(t *testing.T, result string) {
				assert.Empty(t, result)
			},
			wantErr: errors.New("url: the value is valid URL."),
		},
		{
			name: "validation error by max length",
			data: testData{
				key: "",
				url: mockURL,
			},
			prepareMockFn: func(shortRepoMock *mock_short_repo.MockRepository, appConfigMock *mock_config_repo.MockApplicationRepository) {
				short.SetGenerator(func(length ...int) string {
					return strings.Repeat("test", 256)
				})
				shortRepoMock.EXPECT().Exists(ctx, gomock.Any()).AnyTimes().Return(true, nil)
				appConfigMock.EXPECT().Get(gomock.Any()).Return(appConfig, nil)
			},
			want: func(t *testing.T, result string) {
				assert.Empty(t, result)
			},
			wantErr: errors.New("key: the value is 1 ~ 255 characters."),
		},
		{
			name: "exists other error",
			data: testData{
				key: "",
				url: mockURL,
			},
			prepareMockFn: func(shortRepoMock *mock_short_repo.MockRepository, appConfigMock *mock_config_repo.MockApplicationRepository) {
				shortRepoMock.EXPECT().Exists(gomock.Any(), gomock.Any()).Return(false, errors.New("Exists other error."))
				appConfigMock.EXPECT().Get(gomock.Any()).Return(appConfig, nil)
			},
			want: func(t *testing.T, result string) {
				assert.Empty(t, result)
			},
			wantErr: errors.New("Service shortURL: An error occurred while checking for duplicates.: Exists other error."),
		},
		{
			name: "exists other error by loop",
			data: testData{
				key: "",
				url: mockURL,
			},
			prepareMockFn: func(shortRepoMock *mock_short_repo.MockRepository, appConfigMock *mock_config_repo.MockApplicationRepository) {
				var cnt int
				shortRepoMock.EXPECT().Exists(gomock.Any(), gomock.Any()).AnyTimes().DoAndReturn(func(_ context.Context, _ string) (bool, error) {
					if cnt == 1 {
						return false, errors.New("Exists other error.")
					}
					cnt++
					return true, nil
				})
				appConfigMock.EXPECT().Get(gomock.Any()).Return(appConfig, nil)
			},
			want: func(t *testing.T, result string) {
				assert.Empty(t, result)
			},
			wantErr: errors.New("Service shortURL: An error occurred while checking for duplicates.: Exists other error."),
		},
		{
			name: "duplicate error with non key generated",
			data: testData{
				key: mockKey,
				url: mockURL,
			},
			prepareMockFn: func(shortRepoMock *mock_short_repo.MockRepository, appConfigMock *mock_config_repo.MockApplicationRepository) {
				shortRepoMock.EXPECT().Exists(ctx, gomock.Any()).Return(false, nil)
				appConfigMock.EXPECT().Get(gomock.Any()).Return(appConfig, nil)
			},
			want: func(t *testing.T, result string) {
				assert.Empty(t, result)
			},
			wantErr: errors.New("Cannot be used because a duplicate key is specified."),
		},
		{
			name: "put other error",
			data: testData{
				key: "",
				url: mockURL,
			},
			prepareMockFn: func(shortRepoMock *mock_short_repo.MockRepository, appConfigMock *mock_config_repo.MockApplicationRepository) {
				shortRepoMock.EXPECT().Exists(ctx, gomock.Any()).Return(false, nil)
				shortRepoMock.EXPECT().Put(ctx, gomock.Any()).Return(nil, errors.New("Put other error."))
				appConfigMock.EXPECT().Get(gomock.Any()).Return(appConfig, nil)
			},
			want: func(t *testing.T, result string) {
				assert.Empty(t, result)
			},
			wantErr: errors.New("Service shortURL: An error occurred during URL generation.: Put other error."),
		},
	}
	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()
			mockRepo := testMockRepository(mockCtl)
			mockAppCfg := testMockAppRepository(mockCtl)
			tt.prepareMockFn(mockRepo, mockAppCfg)
			appCfg, err := mockAppCfg.Get(ctx)
			require.NoError(t, err)
			beginner, _ := noop.NewBeginner()
			impl := NewService(mockRepo, appCfg, beginner)
			result, err := impl.GenerateShortURL(ctx, tt.data.url, tt.data.key, tt.data.author)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("GenerateShortURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr != nil {
				require.Equal(t, tt.wantErr.Error(), err.Error())
				require.Nil(t, result)
			} else {
				tt.want(t, result.ServiceURL(appCfg.BaseURL))
			}
			short.SetGenerator(nil)
		})
	}
}

func TestServiceForRemove(t *testing.T) {
	var (
		mockKey   = "mock"
		appConfig = &config.Application{
			RetryGenerateCount: 2,
		}
	)
	author := "anonymous"
	hashAuthor := hash.Sum([]byte(author))
	ctx := context.Background()
	tests := []testTable[struct{}, string]{
		{
			name: "success",
			prepareMockFn: func(shortRepoMock *mock_short_repo.MockRepository, appConfigMock *mock_config_repo.MockApplicationRepository) {
				appConfigMock.EXPECT().Get(gomock.Any()).Return(appConfig, nil)
				shortRepoMock.EXPECT().Del(ctx, mockKey, hashAuthor).Return(true, nil)
			},
			want: func(t *testing.T, result string) {
				assert.Equal(t, mockKey, result)
			},
			wantErr: nil,
		},
		{
			name: "success for not found",
			prepareMockFn: func(shortRepoMock *mock_short_repo.MockRepository, appConfigMock *mock_config_repo.MockApplicationRepository) {
				appConfigMock.EXPECT().Get(gomock.Any()).Return(appConfig, nil)
				shortRepoMock.EXPECT().Del(ctx, mockKey, hashAuthor).Return(false, repository.ErrRecordNotFound)
			},
			wantErr: nil,
		},
		{
			name: "internal error",
			prepareMockFn: func(shortRepoMock *mock_short_repo.MockRepository, appConfigMock *mock_config_repo.MockApplicationRepository) {
				appConfigMock.EXPECT().Get(gomock.Any()).Return(appConfig, nil)
				shortRepoMock.EXPECT().Del(ctx, mockKey, hashAuthor).Return(false, errors.New("Error."))
			},
			wantErr: errors.New("Service shortURL: An error occurred during deletion.: Error."),
		},
	}
	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()
			mockRepo := testMockRepository(mockCtl)
			mockAppCfg := testMockAppRepository(mockCtl)
			tt.prepareMockFn(mockRepo, mockAppCfg)
			appCfg, err := mockAppCfg.Get(ctx)
			require.NoError(t, err)
			beginner, _ := noop.NewBeginner()
			impl := NewService(mockRepo, appCfg, beginner)
			err = impl.Remove(ctx, mockKey, "anonymous")
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("GenerateShortURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr != nil {
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
		})
	}
}

func TestServiceForFindAll(t *testing.T) {
	var (
		mockURL   = "http://localhost:8080/example"
		appConfig = &config.Application{
			RetryGenerateCount: 2,
		}
	)
	author := "anonymous"
	hashAuthor := hash.Sum([]byte(author))
	ti, _ := time.Parse("2006-01-02 15:04:05", "2023-06-04 10:10:00")
	ctx := context.Background()
	tests := []testTable[string, []short.ShortWithTimeStamp]{
		{
			name: "success",
			data: author,
			prepareMockFn: func(shortRepoMock *mock_short_repo.MockRepository, appConfigMock *mock_config_repo.MockApplicationRepository) {
				shortRepoMock.EXPECT().FindAll(ctx, hashAuthor).Return([]short.ShortWithTimeStamp{
					{
						Short:     short.NewShort(mockURL, "key", hashAuthor),
						CreatedAt: ti,
						UpdatedAt: ti,
					},
					{
						Short:     short.NewShort(mockURL, "key2", hashAuthor),
						CreatedAt: ti,
						UpdatedAt: ti,
					},
				}, nil)
				appConfigMock.EXPECT().Get(gomock.Any()).Return(appConfig, nil)
			},
			want: func(t *testing.T, result []short.ShortWithTimeStamp) {
				require := require.New(t)
				require.Len(result, 2)
				require.Equal(result[0].GetKey(), "key")
				require.Equal(result[1].GetKey(), "key2")
			},
			wantErr: nil,
		},
		{
			name: "not found",
			data: "not found",
			prepareMockFn: func(shortRepoMock *mock_short_repo.MockRepository, appConfigMock *mock_config_repo.MockApplicationRepository) {
				shortRepoMock.EXPECT().FindAll(ctx, gomock.Any()).Return(nil, repository.ErrRecordNotFound)
				appConfigMock.EXPECT().Get(gomock.Any()).Return(appConfig, nil)
			},
			want: func(t *testing.T, result []short.ShortWithTimeStamp) {
				require := require.New(t)
				require.Len(result, 0)
			},
			wantErr: service.ErrNotFound,
		},
		{
			name: "fail",
			data: "fail",
			prepareMockFn: func(shortRepoMock *mock_short_repo.MockRepository, appConfigMock *mock_config_repo.MockApplicationRepository) {
				shortRepoMock.EXPECT().FindAll(ctx, gomock.Any()).Return(nil, errors.New("Other error"))
				appConfigMock.EXPECT().Get(gomock.Any()).Return(appConfig, nil)
			},
			want: func(t *testing.T, result []short.ShortWithTimeStamp) {
				assert.Empty(t, result)
			},
			wantErr: errors.New("Service shortURL: An error occurred while retrieving data.: Other error"),
		},
	}
	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()
			mockRepo := testMockRepository(mockCtl)
			mockAppCfg := testMockAppRepository(mockCtl)
			tt.prepareMockFn(mockRepo, mockAppCfg)
			appCfg, err := mockAppCfg.Get(ctx)
			require.NoError(t, err)
			beginner, _ := noop.NewBeginner()
			impl := NewService(mockRepo, appCfg, beginner)
			result, err := impl.FindAll(ctx, tt.data)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr != nil {
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
			tt.want(t, result)
		})
	}
}

func TestServiceForQRCodeGenerator(t *testing.T) {
	var (
		mockURL   = "http://localhost:8080/example"
		appConfig = &config.Application{
			RetryGenerateCount: 2,
		}
	)
	fixtureDir, err := fixtures.GetDirectory()
	require.NoError(t, err)
	testQRCode := filepath.Join(fixtureDir, "qrcode.png")
	qrcodeBuf, err := os.ReadFile(testQRCode)
	require.NoError(t, err)
	ctx := context.Background()
	tests := []testTable[string, []byte]{
		{
			name: "success",
			data: "key",
			prepareMockFn: func(shortRepoMock *mock_short_repo.MockRepository, appConfigMock *mock_config_repo.MockApplicationRepository) {
				model := short.NewShort(mockURL, "key", "test")
				shortRepoMock.EXPECT().Get(ctx, "key").Return(model, nil)
				appConfigMock.EXPECT().Get(gomock.Any()).Return(appConfig, nil)
			},
			want: func(t *testing.T, result []byte) {
				assert.Equal(t, qrcodeBuf, result)
			},
			wantErr: nil,
		},
		{
			name: "not found",
			data: "not found",
			prepareMockFn: func(shortRepoMock *mock_short_repo.MockRepository, appConfigMock *mock_config_repo.MockApplicationRepository) {
				shortRepoMock.EXPECT().Get(ctx, gomock.Any()).Return(nil, repository.ErrRecordNotFound)
				appConfigMock.EXPECT().Get(gomock.Any()).Return(appConfig, nil)
			},
			want: func(t *testing.T, result []byte) {
				assert.Empty(t, result)
			},
			wantErr: service.ErrNotFound,
		},
		{
			name: "fail",
			data: "fail",
			prepareMockFn: func(shortRepoMock *mock_short_repo.MockRepository, appConfigMock *mock_config_repo.MockApplicationRepository) {
				shortRepoMock.EXPECT().Get(ctx, gomock.Any()).Return(nil, errors.New("Other error"))
				appConfigMock.EXPECT().Get(gomock.Any()).Return(appConfig, nil)
			},
			want: func(t *testing.T, result []byte) {
				assert.Empty(t, result)
			},
			wantErr: errors.New("Service shortURL: An error occurred while retrieving the URL.: Other error"),
		},
	}
	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()
			mockRepo := testMockRepository(mockCtl)
			mockAppCfg := testMockAppRepository(mockCtl)
			tt.prepareMockFn(mockRepo, mockAppCfg)
			appCfg, err := mockAppCfg.Get(ctx)
			require.NoError(t, err)
			beginner, _ := noop.NewBeginner()
			impl := NewService(mockRepo, appCfg, beginner)
			r, err := impl.GenerateQRCode(ctx, tt.data)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("GenerateQRCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			var buf []byte
			if err != nil && tt.wantErr != nil {
				require.Equal(t, tt.wantErr.Error(), err.Error())
			} else {
				buf, _ = io.ReadAll(r)
			}
			tt.want(t, buf)
		})
	}
}
