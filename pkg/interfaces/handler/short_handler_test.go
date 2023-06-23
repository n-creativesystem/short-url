package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/golang/mock/gomock"
	"github.com/n-creativesystem/short-url/pkg/interfaces/request"
	mock_short "github.com/n-creativesystem/short-url/pkg/mock/service/short"
	"github.com/n-creativesystem/short-url/pkg/service"
	"github.com/n-creativesystem/short-url/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShortHandlerForRedirect(t *testing.T) {
	type testTable struct {
		name           string
		data           string
		prepareMockFn  func(mockSvc *mock_short.MockService)
		prepareRequest func(*http.Request)
		wantLocation   string
	}
	tests := []testTable{
		{
			name: "success",
			data: "aaa",
			prepareMockFn: func(mockSvc *mock_short.MockService) {
				mockSvc.EXPECT().GetURL(gomock.Any(), "aaa").Return("http://localhost:8080/success", nil)
			},
			prepareRequest: func(r *http.Request) {},
			wantLocation:   "http://localhost:8080/success",
		},
		{
			name: "failed",
			data: "abc",
			prepareMockFn: func(mockSvc *mock_short.MockService) {
				mockSvc.EXPECT().GetURL(gomock.Any(), "abc").Return("", errors.New("Not found"))
			},
			prepareRequest: func(r *http.Request) {},
			wantLocation:   "http://localhost/notfound",
		},
		{
			name: "failed for space",
			data: " ",
			prepareMockFn: func(mockSvc *mock_short.MockService) {
			},
			prepareRequest: func(r *http.Request) {},
			wantLocation:   "http://localhost/notfound",
		},
	}
	gin.SetMode(gin.TestMode)
	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()
			mockSvc := mock_short.NewMockService(mockCtl)
			handler := NewShortHandler(mockSvc, WithBaseURL("http://localhost"))
			tt.prepareMockFn(mockSvc)
			u := utils.MustURL("/", tt.data)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, u, nil)
			if tt.prepareRequest != nil {
				tt.prepareRequest(req)
			}
			_, e := gin.CreateTestContext(w)
			handler.ServiceRouter(e)
			e.ServeHTTP(w, req)
			assert.Equal(t, http.StatusTemporaryRedirect, w.Code)
			assert.Equal(t, tt.wantLocation, w.Header().Get("Location"))
		})
	}
}

func TestShortHandlerForGenerateShortURL(t *testing.T) {
	type result struct {
		code   int
		result any
	}
	type testTable struct {
		name           string
		data           request.GenerateShortURL
		prepareMockFn  func(mockSvc *mock_short.MockService)
		prepareRequest func(*http.Request)
		want           result
		accessInfo     oauth2.TokenInfo
	}
	tests := []testTable{
		{
			name: "success",
			data: request.GenerateShortURL{
				URL: "http://localhost:8080/success",
			},
			prepareMockFn: func(mockSvc *mock_short.MockService) {
				mockSvc.EXPECT().GenerateShortURL(gomock.Any(), "http://localhost:8080/success", "", "anonymous").Return("ABC", nil)
			},
			prepareRequest: func(r *http.Request) {},
			want: result{
				code: http.StatusOK,
				result: map[string]interface{}{
					"url": "http://localhost/ABC",
				},
			},
			accessInfo: &models.Token{
				UserID: "anonymous",
			},
		},
		{
			name: "failed json parse",
			data: request.GenerateShortURL{
				URL: "http://localhost:8080/success",
			},
			prepareMockFn: func(mockSvc *mock_short.MockService) {
			},
			prepareRequest: func(r *http.Request) {
				buf := strings.NewReader("aaa=abc")
				r.Body = io.NopCloser(buf)
			},
			want: result{
				code: http.StatusBadRequest,
				result: map[string]interface{}{
					"errors": []interface{}{
						map[string]interface{}{
							"message": "Invalid request",
							"field":   "",
							"help":    "",
						},
					},
				},
			},
			accessInfo: &models.Token{
				UserID: "anonymous",
			},
		},
		{
			name: "failed for valid",
			data: request.GenerateShortURL{
				URL: "",
			},
			prepareMockFn: func(mockSvc *mock_short.MockService) {
			},
			prepareRequest: func(r *http.Request) {},
			want: result{
				code: http.StatusBadRequest,
				result: map[string]interface{}{
					"errors": []interface{}{
						map[string]interface{}{
							"message": "URL is a required field.",
							"field":   "url",
							"help":    "",
						},
					},
				},
			},
			accessInfo: &models.Token{
				UserID: "anonymous",
			},
		},
		{
			name: "failed for service",
			data: request.GenerateShortURL{
				URL: "http://localhost/failed",
			},
			prepareMockFn: func(mockSvc *mock_short.MockService) {
				mockSvc.EXPECT().GenerateShortURL(gomock.Any(), "http://localhost/failed", "", "anonymous").Return("", errors.New("Generate error"))
			},
			prepareRequest: func(r *http.Request) {},
			want: result{
				code: http.StatusInternalServerError,
				result: map[string]interface{}{
					"errors": []interface{}{
						map[string]interface{}{
							"message": "An error occurred during URL generation.",
							"field":   "",
							"help":    "",
						},
					},
				},
			},
			accessInfo: &models.Token{
				UserID: "anonymous",
			},
		},
		{
			name: "failed for unauthorize",
			data: request.GenerateShortURL{
				URL: "http://localhost/failed",
			},
			prepareMockFn: func(mockSvc *mock_short.MockService) {
			},
			prepareRequest: func(r *http.Request) {},
			want: result{
				code: http.StatusUnauthorized,
				result: map[string]interface{}{
					"errors": []interface{}{
						map[string]interface{}{
							"message": http.StatusText(http.StatusUnauthorized),
							"field":   "",
							"help":    "",
						},
					},
				},
			},
		},
	}
	gin.SetMode(gin.TestMode)
	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()
			mockSvc := mock_short.NewMockService(mockCtl)
			handler := NewShortHandler(mockSvc, WithBaseURL("http://localhost"))
			tt.prepareMockFn(mockSvc)
			w := httptest.NewRecorder()
			buf := new(bytes.Buffer)
			err := json.NewEncoder(buf).Encode(&tt.data)
			require.NoError(t, err)
			req, _ := http.NewRequest(http.MethodPost, "/", buf)
			req.Header.Set("Content-Type", "application/json")
			if tt.prepareRequest != nil {
				tt.prepareRequest(req)
			}
			c, _ := gin.CreateTestContext(w)
			c.Request = req
			c.Set(accessInfo, tt.accessInfo)
			handler.GenerateShortURL(c)
			assert.Equal(t, tt.want.code, w.Code)
			var mp any
			err = json.NewDecoder(w.Body).Decode(&mp)
			require.NoError(t, err)
			assert.Contains(t, w.Header().Get("Content-Type"), "application/json")
			assert.Equal(t, tt.want.result, mp)
		})
	}
}

func TestShortHandlerForRemove(t *testing.T) {
	type testTable struct {
		name           string
		data           string
		prepareMockFn  func(mockSvc *mock_short.MockService)
		prepareRequest func(*http.Request)
		wantCode       int
		accessInfo     oauth2.TokenInfo
	}
	tests := []testTable{
		{
			name: "success",
			data: "aaa",
			prepareMockFn: func(mockSvc *mock_short.MockService) {
				mockSvc.EXPECT().Remove(gomock.Any(), "aaa", "anonymous").Return(nil)
			},
			prepareRequest: func(r *http.Request) {},
			wantCode:       http.StatusNoContent,
			accessInfo: &models.Token{
				UserID: "anonymous",
			},
		},
		{
			name: "success(not found)",
			data: "aaa",
			prepareMockFn: func(mockSvc *mock_short.MockService) {
				mockSvc.EXPECT().Remove(gomock.Any(), "aaa", "anonymous").Return(service.ErrNotFound)
			},
			prepareRequest: func(r *http.Request) {},
			wantCode:       http.StatusNoContent,
			accessInfo: &models.Token{
				UserID: "anonymous",
			},
		},
		{
			name: "failed required",
			data: " ",
			prepareMockFn: func(mockSvc *mock_short.MockService) {
			},
			prepareRequest: func(r *http.Request) {},
			wantCode:       http.StatusBadRequest,
			accessInfo: &models.Token{
				UserID: "anonymous",
			},
		},
		{
			name: "failed for service",
			data: "aaa",
			prepareMockFn: func(mockSvc *mock_short.MockService) {
				mockSvc.EXPECT().Remove(gomock.Any(), "aaa", "anonymous").Return(errors.New("Service error"))
			},
			prepareRequest: func(r *http.Request) {},
			wantCode:       http.StatusInternalServerError,
			accessInfo: &models.Token{
				UserID: "anonymous",
			},
		},
		{
			name: "failed for unauthorize",
			data: "aaa",
			prepareMockFn: func(mockSvc *mock_short.MockService) {
			},
			prepareRequest: func(r *http.Request) {},
			wantCode:       http.StatusUnauthorized,
		},
	}
	gin.SetMode(gin.TestMode)
	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()
			mockSvc := mock_short.NewMockService(mockCtl)
			handler := NewShortHandler(mockSvc, WithBaseURL("http://localhost"))
			tt.prepareMockFn(mockSvc)
			u := utils.MustURL("/shorts", tt.data)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodDelete, u, nil)
			if tt.prepareRequest != nil {
				tt.prepareRequest(req)
			}
			_, e := gin.CreateTestContext(w)
			handler.APIRouter(e, func(ctx *gin.Context) {
				ctx.Set(accessInfo, tt.accessInfo)
			})
			e.ServeHTTP(w, req)
			assert.Equal(t, tt.wantCode, w.Code)
		})
	}
}
