package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/golang/mock/gomock"
	"github.com/n-creativesystem/short-url/pkg/domain/config"
	mock_oauth2 "github.com/n-creativesystem/short-url/pkg/mock/external/oauth2"
	mock_oauth2client "github.com/n-creativesystem/short-url/pkg/mock/service/oauth2client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// func TestOAuth2HandlerForRegister(t *testing.T) {
// 	type result struct {
// 		code   int
// 		result any
// 	}
// 	type testTable struct {
// 		name           string
// 		config         *config.Application
// 		prepareMockFn  func(mockSvc *mock_oauth2client.MockService, mockToken *mock_oauth2.MockTokenStore)
// 		prepareRequest func(r *http.Request)
// 		want           result
// 	}
// 	testsTable := []testTable{
// 		{
// 			name: "success generate",
// 			want: result{
// 				code: 200,
// 				result: map[string]interface{}{
// 					"client_id":     "client_id",
// 					"client_secret": "client_secret",
// 				},
// 			},
// 			prepareMockFn: func(mockSvc *mock_oauth2client.MockService, mockToken *mock_oauth2.MockTokenStore) {
// 				result := oauth2client.RegisterResult{
// 					ClientId:     "client_id",
// 					ClientSecret: "client_secret",
// 				}
// 				mockSvc.EXPECT().RegisterClient(gomock.Any(), "anonymous", "app_name").Return(result, nil)
// 			},
// 		},
// 		{
// 			name: "success generate, forwarded",
// 			config: &config.Application{
// 				ForwardedName: "X-User",
// 			},
// 			want: result{
// 				code: 200,
// 				result: map[string]interface{}{
// 					"client_id":     "client_id",
// 					"client_secret": "client_secret",
// 				},
// 			},
// 			prepareMockFn: func(mockSvc *mock_oauth2client.MockService, mockToken *mock_oauth2.MockTokenStore) {
// 				result := oauth2client.RegisterResult{
// 					ClientId:     "client_id",
// 					ClientSecret: "client_secret",
// 				}
// 				mockSvc.EXPECT().RegisterClient(gomock.Any(), "test", "app_name").Return(result, nil)
// 			},
// 			prepareRequest: func(r *http.Request) {
// 				r.Header.Set("X-User", "test")
// 			},
// 		},
// 		{
// 			name: "abort generate",
// 			want: result{
// 				code: 500,
// 				result: map[string]interface{}{
// 					"errors": []interface{}{
// 						map[string]interface{}{
// 							"message": "generate error",
// 							"field":   "",
// 							"help":    "",
// 						},
// 					},
// 				},
// 			},
// 			prepareMockFn: func(mockSvc *mock_oauth2client.MockService, mockToken *mock_oauth2.MockTokenStore) {
// 				result := oauth2client.RegisterResult{
// 					ClientId:     "",
// 					ClientSecret: "",
// 				}
// 				mockSvc.EXPECT().RegisterClient(gomock.Any(), "anonymous", "app_name").Return(result, errors.New("generate error"))
// 			},
// 		},
// 	}
// 	gin.SetMode(gin.TestMode)
// 	t.Parallel()
// 	for _, tt := range testsTable {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			mockCtl := gomock.NewController(t)
// 			defer mockCtl.Finish()
// 			mockSvc := mock_oauth2client.NewMockService(mockCtl)
// 			mockToken := mock_oauth2.NewMockTokenStore(mockCtl)
// 			tt.prepareMockFn(mockSvc, mockToken)
// 			handler := NewOAuthHandler(mockSvc, mockToken, tt.config)
// 			w := httptest.NewRecorder()
// 			req, _ := http.NewRequest(http.MethodPost, "/", nil)
// 			if tt.prepareRequest != nil {
// 				tt.prepareRequest(req)
// 			}
// 			c, _ := gin.CreateTestContext(w)
// 			c.Request = req
// 			handler.RegisterApplication(c)
// 			assert.Equal(t, tt.want.code, w.Code)
// 			var mp any
// 			err := json.NewDecoder(w.Body).Decode(&mp)
// 			require.NoError(t, err)
// 			assert.Equal(t, tt.want.result, mp)
// 		})
// 	}
// }

func TestOAuth2HandlerForTokenRequest(t *testing.T) {
	type testTable struct {
		name           string
		prepareMockFn  func(mockSvc *mock_oauth2client.MockService, mockToken *mock_oauth2.MockTokenStore)
		prepareRequest func(r *http.Request)
		want           int
	}
	testsTable := []testTable{
		{
			name: "success client_credential",
			want: http.StatusOK,
			prepareMockFn: func(mockSvc *mock_oauth2client.MockService, mockToken *mock_oauth2.MockTokenStore) {
				clientInfo := &models.Client{
					ID:     "client_id",
					Secret: "client_secret",
					Domain: "http://localhost",
					Public: false,
					UserID: "anonymous",
				}
				mockSvc.EXPECT().GetByID(gomock.Any(), "client_id").Return(clientInfo, nil)
				mockToken.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
			},
			prepareRequest: func(r *http.Request) {
				value := url.Values{
					"client_id":     {"client_id"},
					"client_secret": {"client_secret"},
					"grant_type":    {"client_credentials"},
				}
				r.Body = io.NopCloser(strings.NewReader(value.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			},
		},
		{
			name: "success refresh_token",
			want: http.StatusOK,
			prepareMockFn: func(mockSvc *mock_oauth2client.MockService, mockToken *mock_oauth2.MockTokenStore) {
				token := models.NewToken()
				token.ClientID = "client_id"
				token.UserID = "anonymous"
				token.Refresh = "old_refresh_token"
				token.Access = "old_access_token"
				mockToken.EXPECT().GetByRefresh(gomock.Any(), "old_refresh_token").Return(token, nil)

				clientInfo := &models.Client{
					ID:     "client_id",
					Secret: "client_secret",
					Domain: "http://localhost",
					Public: false,
					UserID: "anonymous",
				}
				mockSvc.EXPECT().GetByID(gomock.Any(), "client_id").Return(clientInfo, nil)

				mockToken.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
				mockToken.EXPECT().RemoveByAccess(gomock.Any(), "old_access_token").Return(nil)
				mockToken.EXPECT().RemoveByRefresh(gomock.Any(), "old_refresh_token").Return(nil)
			},
			prepareRequest: func(r *http.Request) {
				value := url.Values{
					"client_id":     {"client_id"},
					"client_secret": {"client_secret"},
					"grant_type":    {"refresh_token"},
					"refresh_token": {"old_refresh_token"},
				}
				r.Body = io.NopCloser(strings.NewReader(value.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			},
		},
		{
			name: "failed no support grant_type",
			want: http.StatusUnauthorized,
			prepareMockFn: func(mockSvc *mock_oauth2client.MockService, mockToken *mock_oauth2.MockTokenStore) {
			},
			prepareRequest: func(r *http.Request) {
				value := url.Values{
					"client_id":     {"client_id"},
					"client_secret": {"client_secret"},
					"grant_type":    {"authorization"},
				}
				r.Body = io.NopCloser(strings.NewReader(value.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			},
		},
	}
	gin.SetMode(gin.TestMode)
	t.Parallel()
	for _, tt := range testsTable {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()
			mockSvc := mock_oauth2client.NewMockService(mockCtl)
			mockToken := mock_oauth2.NewMockTokenStore(mockCtl)
			tt.prepareMockFn(mockSvc, mockToken)
			handler := NewOAuthHandler(mockSvc, mockToken, &config.Application{})
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/", nil)
			if tt.prepareRequest != nil {
				tt.prepareRequest(req)
			}
			c, _ := gin.CreateTestContext(w)
			c.Request = req
			handler.TokenRequest(c)
			assert.Equal(t, tt.want, w.Code)
			if tt.want == 200 {
				var mp map[string]interface{}
				err := json.NewDecoder(w.Body).Decode(&mp)
				require.NoError(t, err)
				assert.NotEmpty(t, mp["access_token"])
				assert.NotEmpty(t, mp["refresh_token"])
			}
		})
	}
}

func TestOAuth2HandlerForRevokeRequest(t *testing.T) {
	type testTable struct {
		name           string
		bearerToken    string
		prepareMockFn  func(mockSvc *mock_oauth2client.MockService, mockToken *mock_oauth2.MockTokenStore)
		prepareRequest func(r *http.Request)
		want           int
	}
	testsTable := []testTable{
		{
			name:        "success on revoke access_token",
			want:        http.StatusOK,
			bearerToken: "access_token",
			prepareMockFn: func(mockSvc *mock_oauth2client.MockService, mockToken *mock_oauth2.MockTokenStore) {
				token := models.NewToken()
				token.ClientID = "client_id"
				token.UserID = "anonymous"
				token.Refresh = "refresh_token"
				token.Access = "access_token"
				mockToken.EXPECT().GetByAccess(gomock.Any(), "access_token").Return(token, nil).MaxTimes(2)
				mockToken.EXPECT().RemoveByAccess(gomock.Any(), "access_token").Return(nil)
			},
			prepareRequest: func(r *http.Request) {
				value := url.Values{
					"token":           {"access_token"},
					"token_type_hint": {"access_token"},
				}
				r.Body = io.NopCloser(strings.NewReader(value.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			},
		},
		{
			name:        "success on revoke refresh_token",
			want:        http.StatusOK,
			bearerToken: "access_token",
			prepareMockFn: func(mockSvc *mock_oauth2client.MockService, mockToken *mock_oauth2.MockTokenStore) {
				token := models.NewToken()
				token.ClientID = "client_id"
				token.UserID = "anonymous"
				token.Refresh = "refresh_token"
				token.Access = "access_token"
				mockToken.EXPECT().GetByAccess(gomock.Any(), "access_token").Return(token, nil)
				mockToken.EXPECT().GetByRefresh(gomock.Any(), "refresh_token").Return(token, nil)
				mockToken.EXPECT().RemoveByRefresh(gomock.Any(), "refresh_token").Return(nil)
			},
			prepareRequest: func(r *http.Request) {
				value := url.Values{
					"token":           {"refresh_token"},
					"token_type_hint": {"refresh_token"},
				}
				r.Body = io.NopCloser(strings.NewReader(value.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			},
		},
		{
			name:        "success on revoke access_token with not token_type_hint",
			want:        http.StatusOK,
			bearerToken: "access_token",
			prepareMockFn: func(mockSvc *mock_oauth2client.MockService, mockToken *mock_oauth2.MockTokenStore) {
				token := models.NewToken()
				token.ClientID = "client_id"
				token.UserID = "anonymous"
				token.Refresh = "refresh_token"
				token.Access = "access_token"
				mockToken.EXPECT().GetByAccess(gomock.Any(), "access_token").Return(token, nil).MaxTimes(2)
				mockToken.EXPECT().RemoveByAccess(gomock.Any(), "access_token").Return(nil)
			},
			prepareRequest: func(r *http.Request) {
				value := url.Values{
					"token": {"access_token"},
				}
				r.Body = io.NopCloser(strings.NewReader(value.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			},
		},
		{
			name:        "success on revoke refresh_token with not token_type_hint",
			want:        http.StatusOK,
			bearerToken: "access_token",
			prepareMockFn: func(mockSvc *mock_oauth2client.MockService, mockToken *mock_oauth2.MockTokenStore) {
				token := models.NewToken()
				token.ClientID = "client_id"
				token.UserID = "anonymous"
				token.Refresh = "refresh_token"
				token.Access = "access_token"
				count := 0
				mockToken.EXPECT().GetByAccess(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, access_token string) (oauth2.TokenInfo, error) {
					count++
					if count == 1 {
						return token, nil
					}
					return nil, nil
				}).MaxTimes(2)
				mockToken.EXPECT().GetByRefresh(gomock.Any(), "refresh_token").Return(token, nil)
				mockToken.EXPECT().RemoveByRefresh(gomock.Any(), "refresh_token").Return(nil)
			},
			prepareRequest: func(r *http.Request) {
				value := url.Values{
					"token": {"refresh_token"},
				}
				r.Body = io.NopCloser(strings.NewReader(value.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			},
		},
		{
			name:        "invalid request for token",
			want:        http.StatusBadRequest,
			bearerToken: "access_token",
			prepareMockFn: func(mockSvc *mock_oauth2client.MockService, mockToken *mock_oauth2.MockTokenStore) {
				token := models.NewToken()
				token.ClientID = "client_id"
				token.UserID = "anonymous"
				token.Refresh = "refresh_token"
				token.Access = "access_token"
				mockToken.EXPECT().GetByAccess(gomock.Any(), "access_token").Return(token, nil)
			},
			prepareRequest: func(r *http.Request) {
				value := url.Values{
					"token": {""},
				}
				r.Body = io.NopCloser(strings.NewReader(value.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			},
		},
		{
			name:        "Error during processing of access_token",
			want:        http.StatusBadRequest,
			bearerToken: "access_token",
			prepareMockFn: func(mockSvc *mock_oauth2client.MockService, mockToken *mock_oauth2.MockTokenStore) {
				token := models.NewToken()
				token.ClientID = "client_id"
				token.UserID = "anonymous"
				token.Refresh = "refresh_token"
				token.Access = "access_token"
				mockToken.EXPECT().GetByAccess(gomock.Any(), "access_token").Return(token, nil)
				mockToken.EXPECT().GetByAccess(gomock.Any(), "access_token").Return(nil, errors.New("Other error"))
			},
			prepareRequest: func(r *http.Request) {
				value := url.Values{
					"token":           {"access_token"},
					"token_type_hint": {"access_token"},
				}
				r.Body = io.NopCloser(strings.NewReader(value.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			},
		},
		{
			name:        "Error during processing of access_token2",
			want:        http.StatusBadRequest,
			bearerToken: "access_token",
			prepareMockFn: func(mockSvc *mock_oauth2client.MockService, mockToken *mock_oauth2.MockTokenStore) {
				token := models.NewToken()
				token.ClientID = "client_id"
				token.UserID = "anonymous"
				token.Refresh = "refresh_token"
				token.Access = "access_token"
				mockToken.EXPECT().GetByAccess(gomock.Any(), "access_token").Return(token, nil).MaxTimes(2)
				mockToken.EXPECT().RemoveByAccess(gomock.Any(), "access_token").Return(errors.New("Other error"))
			},
			prepareRequest: func(r *http.Request) {
				value := url.Values{
					"token":           {"access_token"},
					"token_type_hint": {"access_token"},
				}
				r.Body = io.NopCloser(strings.NewReader(value.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			},
		},
		{
			name:        "Error during processing of access_token3",
			want:        http.StatusBadRequest,
			bearerToken: "access_token",
			prepareMockFn: func(mockSvc *mock_oauth2client.MockService, mockToken *mock_oauth2.MockTokenStore) {
				token := models.NewToken()
				token.ClientID = "client_id"
				token.UserID = "anonymous"
				token.Refresh = "refresh_token"
				token.Access = "access_token"
				mockToken.EXPECT().GetByAccess(gomock.Any(), "access_token").Return(token, nil).MaxTimes(2)
				mockToken.EXPECT().RemoveByAccess(gomock.Any(), "access_token").Return(errors.New("Other error"))
			},
			prepareRequest: func(r *http.Request) {
				value := url.Values{
					"token": {"access_token"},
				}
				r.Body = io.NopCloser(strings.NewReader(value.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			},
		},
		{
			name:        "Error during processing of refresh_token",
			want:        http.StatusBadRequest,
			bearerToken: "access_token",
			prepareMockFn: func(mockSvc *mock_oauth2client.MockService, mockToken *mock_oauth2.MockTokenStore) {
				token := models.NewToken()
				token.ClientID = "client_id"
				token.UserID = "anonymous"
				token.Refresh = "refresh_token"
				token.Access = "access_token"
				mockToken.EXPECT().GetByAccess(gomock.Any(), "access_token").Return(token, nil)
				mockToken.EXPECT().GetByRefresh(gomock.Any(), "refresh_token").Return(nil, errors.New("Other error"))
			},
			prepareRequest: func(r *http.Request) {
				value := url.Values{
					"token":           {"refresh_token"},
					"token_type_hint": {"refresh_token"},
				}
				r.Body = io.NopCloser(strings.NewReader(value.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			},
		},
		{
			name:        "Error during processing of refresh_token2",
			want:        http.StatusBadRequest,
			bearerToken: "access_token",
			prepareMockFn: func(mockSvc *mock_oauth2client.MockService, mockToken *mock_oauth2.MockTokenStore) {
				token := models.NewToken()
				token.ClientID = "client_id"
				token.UserID = "anonymous"
				token.Refresh = "refresh_token"
				token.Access = "access_token"
				mockToken.EXPECT().GetByAccess(gomock.Any(), "access_token").Return(token, nil)
				mockToken.EXPECT().GetByRefresh(gomock.Any(), "refresh_token").Return(nil, nil)
			},
			prepareRequest: func(r *http.Request) {
				value := url.Values{
					"token":           {"refresh_token"},
					"token_type_hint": {"refresh_token"},
				}
				r.Body = io.NopCloser(strings.NewReader(value.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			},
		},
		{
			name:        "Error during processing of refresh_token3",
			want:        http.StatusBadRequest,
			bearerToken: "access_token",
			prepareMockFn: func(mockSvc *mock_oauth2client.MockService, mockToken *mock_oauth2.MockTokenStore) {
				token := models.NewToken()
				token.ClientID = "client_id"
				token.UserID = "anonymous"
				token.Refresh = "refresh_token"
				token.Access = "access_token"
				mockToken.EXPECT().GetByAccess(gomock.Any(), "access_token").Return(token, nil)
				mockToken.EXPECT().GetByRefresh(gomock.Any(), "refresh_token").Return(token, nil)
				mockToken.EXPECT().RemoveByRefresh(gomock.Any(), "refresh_token").Return(errors.New("Other error"))
			},
			prepareRequest: func(r *http.Request) {
				value := url.Values{
					"token":           {"refresh_token"},
					"token_type_hint": {"refresh_token"},
				}
				r.Body = io.NopCloser(strings.NewReader(value.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			},
		},
		{
			name:        "Error during processing of refresh_token3",
			want:        http.StatusBadRequest,
			bearerToken: "access_token",
			prepareMockFn: func(mockSvc *mock_oauth2client.MockService, mockToken *mock_oauth2.MockTokenStore) {
				token := models.NewToken()
				token.ClientID = "client_id"
				token.UserID = "anonymous"
				token.Refresh = "refresh_token"
				token.Access = "access_token"
				mockToken.EXPECT().GetByAccess(gomock.Any(), "access_token").Return(token, nil)
				mockToken.EXPECT().GetByAccess(gomock.Any(), "refresh_token").Return(nil, nil)
				mockToken.EXPECT().GetByRefresh(gomock.Any(), "refresh_token").Return(token, nil)
				mockToken.EXPECT().RemoveByRefresh(gomock.Any(), "refresh_token").Return(errors.New("Other error"))
			},
			prepareRequest: func(r *http.Request) {
				value := url.Values{
					"token": {"refresh_token"},
				}
				r.Body = io.NopCloser(strings.NewReader(value.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			},
		},
		{
			name:        "invalid request for token_type_hint with no supported",
			want:        http.StatusBadRequest,
			bearerToken: "access_token",
			prepareMockFn: func(mockSvc *mock_oauth2client.MockService, mockToken *mock_oauth2.MockTokenStore) {
				token := models.NewToken()
				token.ClientID = "client_id"
				token.UserID = "anonymous"
				token.Refresh = "refresh_token"
				token.Access = "access_token"
				mockToken.EXPECT().GetByAccess(gomock.Any(), "access_token").Return(token, nil)
			},
			prepareRequest: func(r *http.Request) {
				value := url.Values{
					"token":           {"access_token"},
					"token_type_hint": {"no_support"},
				}
				r.Body = io.NopCloser(strings.NewReader(value.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			},
		},
	}
	gin.SetMode(gin.TestMode)
	t.Parallel()
	for _, tt := range testsTable {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()
			mockSvc := mock_oauth2client.NewMockService(mockCtl)
			mockToken := mock_oauth2.NewMockTokenStore(mockCtl)
			tt.prepareMockFn(mockSvc, mockToken)
			handler := NewOAuthHandler(mockSvc, mockToken, &config.Application{})
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/oauth2/revoke", nil)
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tt.bearerToken))
			if tt.prepareRequest != nil {
				tt.prepareRequest(req)
			}
			c, e := gin.CreateTestContext(w)
			handler.Router(e)
			c.Request = req
			e.ServeHTTP(w, req)
			assert.Equal(t, tt.want, w.Code)
		})
	}
}
