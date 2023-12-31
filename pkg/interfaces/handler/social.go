package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sort"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"github.com/n-creativesystem/short-url/pkg/domain/social"
	"github.com/n-creativesystem/short-url/pkg/interfaces/middleware/session"
	"github.com/n-creativesystem/short-url/pkg/interfaces/response"
	"github.com/n-creativesystem/short-url/pkg/utils/compare"
	"github.com/n-creativesystem/short-url/pkg/utils/logging"
	"golang.org/x/oauth2"
)

type Social struct {
	cfg        *social.Config
	socialRepo social.UserRepository
}

func (s *Social) Authorization(ctx context.Context, socialId string) string {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	state := randomString(32)
	nonce := randomString(32)
	authURL := s.cfg.Oauth2Config.AuthCodeURL(state, oidc.Nonce(nonce))
	sm := session.GetContext(ctx)
	sm.Put(ctx, "state", state)
	sm.Put(ctx, "nonce", nonce)
	return authURL
}

type CallbackResult struct {
	Code int
	Err  error
}

func (c *CallbackResult) setError(code int, err error) *CallbackResult {
	c.Code = code
	c.Err = err
	return c
}

func (s *Social) Callback(r *http.Request) *CallbackResult {
	ctx := r.Context()
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	*r = *r.WithContext(ctx)
	result := &CallbackResult{}
	sm := session.GetContext(ctx)
	state := r.URL.Query().Get("state")
	sessionState := sm.PopString(ctx, "state")
	sessionNonce := sm.PopString(ctx, "nonce")
	if !compare.ConstantTimeCompare(state, sessionState) {
		return result.setError(http.StatusBadRequest, errors.New("state validation failed"))
	}
	code := r.URL.Query().Get("code")
	oauth2Token, err := s.cfg.Oauth2Config.Exchange(ctx, code)
	if err != nil {
		logging.FromContext(ctx).With(logging.WithErr(err)).ErrorContext(ctx, fmt.Sprintf("Exchange token: %v", err))
		return result.setError(http.StatusUnauthorized, errors.New("Failed to exchange token"))
	}
	rawIdToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		err := errors.New("Mission id_token")
		logging.FromContext(ctx).With(logging.WithErr(err)).ErrorContext(ctx, fmt.Sprintf("Extra id_token: %v", err))
		return result.setError(http.StatusInternalServerError, err)
	}
	oidcConfig := &oidc.Config{
		ClientID: s.cfg.Oauth2Config.ClientID,
	}
	verify := s.cfg.Provider.Verifier(oidcConfig)
	idToken, err := verify.Verify(ctx, rawIdToken)
	if err != nil {
		logging.FromContext(ctx).With(logging.WithErr(err)).ErrorContext(ctx, fmt.Sprintf("Verify id_token: %v", err))
		return result.setError(http.StatusInternalServerError, err)
	}
	if !compare.ConstantTimeCompare(idToken.Nonce, sessionNonce) {
		err := errors.New("nonce validation failed")
		logging.FromContext(ctx).With(logging.WithErr(err)).ErrorContext(ctx, "nonce validation failed")
		return result.setError(http.StatusInternalServerError, err)
	}
	u, err := s.cfg.Provider.UserInfo(ctx, oauth2.StaticTokenSource(oauth2Token))
	if err != nil {
		logging.FromContext(ctx).With(logging.WithErr(err)).ErrorContext(ctx, "Failed to request of user info")
		return result.setError(http.StatusInternalServerError, errors.New("Failed to request of user info"))
	}
	user := &social.User{UserInfo: u}
	if err := user.ParseClaims(s.cfg.ClaimKeys); err != nil {
		logging.FromContext(ctx).With(logging.WithErr(err)).WarnContext(ctx, err.Error())
	}
	user, err = s.socialRepo.Register(ctx, user)
	if err != nil {
		logging.FromContext(ctx).With(logging.WithErr(err)).ErrorContext(ctx, "Failed to register user")
		return result.setError(http.StatusInternalServerError, errors.New("Failed to register user"))
	}
	sm.Put(ctx, "loginUser", string(user.Encode()))
	return result
}

type SocialHandler struct {
	providers        map[string]Social
	LoginSuccessURL  string
	LogoutSuccessURL string
}

func NewSocialHandler(config map[string]*social.Config, loginSuccessURL, logoutSuccessURL string, socialRepo social.UserRepository) *SocialHandler {
	if loginSuccessURL == "" {
		loginSuccessURL = "/"
	}
	p := map[string]Social{}
	for key, cfg := range config {
		p[key] = Social{
			cfg:        cfg,
			socialRepo: socialRepo,
		}
	}
	return &SocialHandler{
		providers:        p,
		LoginSuccessURL:  loginSuccessURL,
		LogoutSuccessURL: logoutSuccessURL,
	}
}

// Authorization
//
// @Summary 認証開始エンドポイント
// @Tags UI
// @Accept json
// @Produce json
// @Success 302 {} Redirect
// @Header 302 {string} Location /{authURL}
// @Router /auth/{provider}/authorize [get]
// @ID SocialLoginAuthorize
func (h *SocialHandler) Authorization(socialId string) gin.HandlerFunc {
	social := h.providers[socialId]
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx, span := tracer.Start(ctx, "")
		defer span.End()
		*c.Request = *c.Request.WithContext(ctx)
		authURL := social.Authorization(ctx, socialId)
		c.Redirect(http.StatusFound, authURL)
	}
}

// Callback
//
// @Summary コールバック
// @Tags UI
// @Accept json
// @Produce json
// @Success 302 {} Redirect
// @Header 302 {string} Location /{loginSuccessURL}
// @Failure 400 {object} response.Errors
// @Failure 401 {object} response.Errors
// @Failure 500 {object} response.Errors
// @Router /auth/{provider}/callback [get]
// @ID SocialLoginCallback
func (h *SocialHandler) Callback(socialId string) gin.HandlerFunc {
	social := h.providers[socialId]
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx, span := tracer.Start(ctx, "")
		defer span.End()
		*c.Request = *c.Request.WithContext(ctx)
		result := social.Callback(c.Request)
		if result.Err != nil {
			c.AbortWithStatusJSON(result.Code, response.NewErrors(result.Err))
			return
		}
		c.Redirect(http.StatusFound, h.LoginSuccessURL)
	}
}

// UserInfo
//
// @Summary ログイン済みのユーザー情報取得
// @Tags UI
// @Accept json
// @Produce json
// @Success 200 {object} response.User
// @Failure 401 {object} response.Errors
// @Router /auth/userinfo [get]
// @ID SocialLoginUserInfo
func (h *SocialHandler) UserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx, span := tracer.Start(ctx, "")
		defer span.End()
		*c.Request = *c.Request.WithContext(ctx)
		user, ok := session.GetAuthUserWithGinContext(c)
		if !ok {
			return
		}
		response := &response.User{
			UserInfo: user.UserInfo,
			Username: user.Username,
			Picture:  user.Picture,
		}
		c.JSON(http.StatusOK, response)
	}
}

// Logout
//
// @Summary ログアウト
// @Tags UI
// @Accept json
// @Produce json
// @Success 302 {} Redirect
// @Header 302 {string} Location /{logoutSuccessURL}
// @Router /auth/logout [get]
// @ID SocialLogout
func (h *SocialHandler) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx, span := tracer.Start(ctx, "")
		defer span.End()
		*c.Request = *c.Request.WithContext(ctx)
		sm := session.GetContext(ctx)
		_ = sm.Destroy(ctx)
		c.Redirect(http.StatusFound, h.LogoutSuccessURL)
	}
}

// GetEnabledSocialLogin
//
// @Summary 有効なソーシャルログイン
// @Tags UI
// @Accept json
// @Produce json
// @Success 200 {object} response.EnabledSocialLogin
// @Router /auth/enabled [get]
// @ID EnabledSocialLoginList
func (h *SocialHandler) GetEnabledSocialLogin() gin.HandlerFunc {
	providers := make([]string, 0, len(h.providers))
	for key := range h.providers {
		providers = append(providers, key)
	}
	sort.Strings(providers)
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx, span := tracer.Start(ctx, "")
		defer span.End()
		*c.Request = *c.Request.WithContext(ctx)
		response := &response.EnabledSocialLogin{
			Socials: providers,
		}
		c.JSON(http.StatusOK, response)
	}
}

func (h *SocialHandler) Router(router gin.IRouter, protected gin.HandlerFunc) {
	g := router.Group("/auth")
	{
		for key := range h.providers {
			g.GET(fmt.Sprintf("/%s/authorize", key), h.Authorization(key))
			g.GET(fmt.Sprintf("/%s/callback", key), h.Callback(key))
		}
		g.GET("/enabled", h.GetEnabledSocialLogin())
		g.GET("/userinfo", protected, h.UserInfo())
		g.GET("/logout", protected, h.Logout())
	}
}
