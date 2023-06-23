package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	_ "embed"

	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4"
	oauth2_errors "github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	_ "github.com/go-sql-driver/mysql"
	"github.com/n-creativesystem/short-url/pkg/domain/config"
	"github.com/n-creativesystem/short-url/pkg/interfaces/middleware/session"
	"github.com/n-creativesystem/short-url/pkg/interfaces/request"
	"github.com/n-creativesystem/short-url/pkg/interfaces/response"
	"github.com/n-creativesystem/short-url/pkg/service"
	oauth2client "github.com/n-creativesystem/short-url/pkg/service/oauth2_client"
	"github.com/n-creativesystem/short-url/pkg/utils/logging"
)

const (
	accessInfo = "access_info"
)

type OAuthHandler struct {
	srv         *server.Server
	service     oauth2client.Service
	appConfig   *config.Application
	oauth2Store oauth2.TokenStore
}

func NewOAuthHandler(service oauth2client.Service, tokenStore oauth2.TokenStore, appConfig *config.Application) *OAuthHandler {
	manager := manage.NewDefaultManager()
	twoWeekHour := (24 * time.Hour) * 14 // 14日
	manage.DefaultClientTokenCfg.RefreshTokenExp = twoWeekHour
	manage.DefaultClientTokenCfg.IsGenerateRefresh = true
	manage.DefaultRefreshTokenCfg.AccessTokenExp = time.Hour * 2
	manage.DefaultRefreshTokenCfg.RefreshTokenExp = twoWeekHour

	manager.MapTokenStorage(tokenStore)
	manager.MapClientStorage(service)
	srv := server.NewDefaultServer(manager)
	srv.SetAllowedResponseType(oauth2.Token)
	srv.SetAllowedGrantType(oauth2.ClientCredentials, oauth2.Refreshing)
	srv.SetClientInfoHandler(clientInfoHandler())
	srv.SetInternalErrorHandler(func(err error) *oauth2_errors.Response {
		logging.Default().Errorf("Internal error: %v", err)
		return nil
	})
	srv.SetResponseErrorHandler(func(re *oauth2_errors.Response) {
		logging.Default().Errorf("Response Error: %v", re.Error)
		re.StatusCode = http.StatusBadRequest
	})
	srv.ResponseTokenHandler = responseToken

	return &OAuthHandler{
		srv:         srv,
		service:     service,
		appConfig:   appConfig,
		oauth2Store: tokenStore,
	}
}

func (h *OAuthHandler) Router(route gin.IRouter) {
	g := route.Group("/oauth2")
	{
		// g.POST("/register", h.RegisterApplication)
		g.GET("", h.ValidationBearerToken(), h.OAuthApplications)
		g.POST("/token", h.TokenRequest)
		g.POST("/revoke", h.ValidationBearerToken(), h.RevokeTokenRequest)
	}
}

func (h *OAuthHandler) RegisterApplication(c *gin.Context) {
	ctx := c.Request.Context()
	forwarded := "anonymous"
	if h.appConfig != nil && h.appConfig.ForwardedName != "" {
		v := c.GetHeader(h.appConfig.ForwardedName)
		if v != "" {
			forwarded = v
		}
	} else {
		if user, ok := session.GetAuthUserWithGinContext(c); ok {
			forwarded = user.Email
		}
	}
	var registerClient request.RegisterApplication
	if !bindJSON(c, &registerClient) {
		return
	}
	if !validation(c, &registerClient) {
		return
	}
	result, err := h.service.RegisterClient(ctx, forwarded, registerClient.AppName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.NewErrors(err))
		return
	}
	resp := response.RegisterOAuth2Client{
		ClientID:     result.ClientId,
		ClientSecret: result.ClientSecret,
	}
	c.JSON(http.StatusOK, &resp)
}

func (h *OAuthHandler) ValidationBearerToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		info, err := h.srv.ValidationBearerToken(r)
		if err != nil {
			data, _, header := h.srv.GetErrorData(oauth2_errors.ErrAccessDenied)
			body := failureBody(data)
			for key, value := range header {
				c.Header(key, strings.Join(value, ";"))
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, &body)
			return
		}
		c.Set(accessInfo, info)
		c.Next()
	}
}

// TokenRequest
//
// @Summary アクセストークンの生成
// @Tags API
// @Accept x-www-form-urlencoded
// @Produce json
// @Param client_id formData string true "client id"
// @Param client_secret formData string true "client secret"
// @Param grant_type formData string true "grant_type" Enums(client_credentials)
// @Success 200 {object} response.OAuth2Token
// @Failure 400 {object} response.Errors
// @Router /oauth2/token [post]
// @ID OAuthTokenRequest
func (h *OAuthHandler) TokenRequest(c *gin.Context) {
	_ = h.srv.HandleTokenRequest(c.Writer, c.Request)
}

// RevokeTokenRequest トークンの取消
//
// @refs: https://openid-foundation-japan.github.io/rfc7009.ja.html#anchor2
// @Summary アクセストークンの取消
// @Tags API
// @Accept json
// @Produce json
// @Param request_body body request.RevokeToken true "revoke token"
// @Success 200
// @Failure 400 {object} response.Errors
// @Failure 401 {object} response.Errors
// @Router /oauth2/revoke [post]
// @Security OAuth2Application
// @ID RevokeOAuthToken
func (h *OAuthHandler) RevokeTokenRequest(c *gin.Context) {
	ctx := c.Request.Context()
	errFn := func() {
		data, _, header := h.srv.GetErrorData(oauth2_errors.ErrInvalidRequest)
		body := failureBody(data)
		for key, values := range header {
			c.Header(key, strings.Join(values, ";"))
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, &body)
	}
	var reqBody request.RevokeToken
	if !bindJSON(c, &reqBody) {
		errFn()
		return
	}
	token := reqBody.Token
	if token == "" {
		errFn()
		return
	}

	tokenHint := reqBody.TokenTypeHint
	if tokenHint != "" {
		switch tokenHint {
		case "access_token":
			_, err := h.removeAccessToken(ctx, token)
			if err != nil {
				errFn()
				return
			}
		case "refresh_token":
			_, err := h.removeRefreshToken(ctx, token)
			if err != nil {
				errFn()
				return
			}
		default:
			errFn()
			return
		}
	} else {
		found, err := h.removeAccessToken(ctx, token)
		if found {
			if err != nil {
				errFn()
				return
			}
		} else {
			_, err = h.removeRefreshToken(ctx, token)
			if err != nil {
				errFn()
				return
			}
		}
	}
	c.Status(http.StatusOK)
}

// OAuthApplications 登録されているOAuthアプリケーションの一覧
//
// @Summary 登録されているOAuthアプリケーションの一覧
// @Tags API
// @Accept json
// @Produce json
// @Success 200
// @Failure 400 {object} response.Errors
// @Failure 401 {object} response.Errors
// @Router /oauth2 [get]
// @Security OAuth2Application
// @ID OAuthApplicationList
func (h *OAuthHandler) OAuthApplications(c *gin.Context) {
	token, ok := getContext[oauth2.TokenInfo](c, accessInfo)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewErrorsWithMessage(http.StatusText(http.StatusUnauthorized)))
		return
	}
	var result response.OAuth2Applications
	ctx := c.Request.Context()
	values, err := h.service.FindAll(ctx, token.GetUserID())
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			c.JSON(http.StatusOK, result)
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, response.NewErrors(err))
		return
	}
	for _, value := range values {
		app := response.OAuth2Application{
			ClientID: value.GetID(),
			Name:     value.GetAppName(),
		}
		result.Results = append(result.Results, app)
	}
	c.JSON(http.StatusOK, &result)
}

func (h *OAuthHandler) removeAccessToken(ctx context.Context, token string) (bool, error) {
	ti, err := h.oauth2Store.GetByAccess(ctx, token)
	if err != nil {
		return false, err
	}
	if ti == nil {
		return false, oauth2_errors.ErrInvalidRequest
	}
	err = h.srv.Manager.RemoveAccessToken(ctx, token)
	if err != nil {
		return true, err
	}
	return true, nil
}

func (h *OAuthHandler) removeRefreshToken(ctx context.Context, token string) (bool, error) {
	ti, err := h.oauth2Store.GetByRefresh(ctx, token)
	if err != nil {
		return false, err
	}
	if ti == nil {
		return false, oauth2_errors.ErrInvalidRequest
	}
	err = h.srv.Manager.RemoveRefreshToken(ctx, token)
	if err != nil {
		return true, err
	}
	return true, nil
}

func clientInfoHandler() server.ClientInfoHandler {
	return func(r *http.Request) (string, string, error) {
		clientID, clientSecret, err := server.ClientFormHandler(r)
		if err != nil {
			return server.ClientBasicHandler(r)
		}
		return clientID, clientSecret, err
	}
}

func responseToken(w http.ResponseWriter, data map[string]interface{}, header http.Header, statusCode ...int) error {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Pragma", "no-cache")

	for key := range header {
		w.Header().Set(key, header.Get(key))
	}

	status := http.StatusOK
	if len(statusCode) > 0 && statusCode[0] > 0 {
		status = statusCode[0]
	}
	var respBody interface{}
	switch status {
	case http.StatusOK:
		respBody = successBody(data)
	case http.StatusBadRequest:
		respBody = failureBody(data)
	default:
		respBody = data
	}

	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(&respBody)
}

func successBody(data map[string]interface{}) response.OAuth2Token {
	var body response.OAuth2Token
	if v, ok := data["access_token"].(string); ok {
		body.AccessToken = v
	}
	if v, ok := data["token_type"].(string); ok {
		body.TokenType = v
	}
	if v, ok := data["expires_in"].(int64); ok {
		body.ExpiresIn = v
	}
	if v, ok := data["scope"].(string); ok {
		body.Scope = v
	}
	if v, ok := data["refresh_token"].(string); ok {
		body.RefreshToken = v
	}
	return body
}

func failureBody(data map[string]interface{}) response.Errors {
	var (
		body response.Errors
		err  response.Error
	)
	if v, ok := data["error"].(string); ok {
		err.Message = v
	}
	if v, ok := data["error_description"].(string); ok {
		err.Description = v
	}
	body.Add(err)
	return body
}

// NOTE: oauth2のライブラリがjsonをサポートしたら下記をコメントインするがoauthのトークン生成はform送信が推奨されている
// @refs: https://www.rfc-editor.org/rfc/rfc6749#section-4.1.3
/*
func clientInfoHandler(r *http.Request) (string, string, error) {
	contentType := convertContentTypeToEnumWithRequest(r)
	switch contentType {
	case jsonType:
		defer r.Body.Close()
		return clientJSONHandler(r)
	case formType:
		return server.ClientFormHandler(r)
	default:
		return "", "", errors.New("No support content-type")
	}
}

func clientJSONHandler(r *http.Request) (string, string, error) {
	var body struct {
		ClientId     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return "", "", err
	}
	return body.ClientId, body.ClientSecret, nil
}
*/
