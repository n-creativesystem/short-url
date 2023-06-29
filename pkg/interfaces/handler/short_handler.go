package handler

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/n-creativesystem/short-url/pkg/interfaces/request"
	"github.com/n-creativesystem/short-url/pkg/interfaces/response"
	"github.com/n-creativesystem/short-url/pkg/service"
	"github.com/n-creativesystem/short-url/pkg/service/short"
	"github.com/n-creativesystem/short-url/pkg/utils"
	"github.com/n-creativesystem/short-url/pkg/utils/logging"
	"github.com/n-creativesystem/short-url/pkg/utils/types"
)

type ShortURLHandler struct {
	service short.Service
	option  shortOption
}

func NewShortHandler(service short.Service, opts ...ShortOption) *ShortURLHandler {
	opt := newShortOption()
	for _, o := range opts {
		o.apply(&opt)
	}

	return &ShortURLHandler{
		service: service,
		option:  opt,
	}
}

func (h *ShortURLHandler) Redirect(c *gin.Context) {
	key := h.getKey(c)
	v, err := key.Value()
	if err != nil {
		h.redirect(c, err)
		return
	}
	ctx := c.Request.Context()
	url, err := h.service.GetURL(ctx, v)
	if err != nil {
		h.redirect(c, err)
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// GenerateShortURL 短縮URLの生成
//
// @Summary 短縮URLの生成
// @Tags API
// @Produce json
// @Param request_body body request.GenerateShortURL true "generate short url request"
// @Success 200 {object} response.GenerateShortURL
// @Failure 400 {object} response.Errors
// @Failure 401 {object} response.Errors
// @Failure 500 {object} response.Errors
// @Router /shorts/generate [post]
// @Security OAuth2Application
// @ID GenerateShortURL
func (h *ShortURLHandler) GenerateShortURL(c *gin.Context) {
	token, ok := getContext[oauth2.TokenInfo](c, accessInfo)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewErrorsWithMessage(http.StatusText(http.StatusUnauthorized)))
		return
	}
	var req request.GenerateShortURL
	if !bindJSON(c, &req) {
		return
	}
	if !validation(c, &req) {
		return
	}
	ctx := c.Request.Context()
	result, err := h.service.GenerateShortURL(ctx, req.URL, req.Key, token.GetUserID())
	if err != nil {
		logging.Default().Error(err)
		var clientErr *service.ClientError
		if errors.As(err, &clientErr) {
			errRes := response.NewErrorsWithMessage(clientErr.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, &errRes)
			return
		}
		errRes := response.NewErrorsWithMessage("An error occurred during URL generation.")
		c.AbortWithStatusJSON(http.StatusInternalServerError, errRes)
		return
	}
	res := &response.GenerateShortURL{
		URL: result.ServiceURL(h.option.appConfig.BaseURL),
	}
	c.JSON(http.StatusOK, res)
}

// GenerateQRCode QRコードの生成
//
// @Summary 短縮URLのQRコード生成
// @Tags API
// @Produce png
// @Param key path string false "short url request"
// @Success 200
// @Failure 400 {object} response.Errors
// @Failure 401 {object} response.Errors
// @Failure 500 {object} response.Errors
// @Router /shorts/{key}/qrcode [get]
// @Security OAuth2Application
// @ID GenerateQRCode
func (h *ShortURLHandler) GenerateQRCode(c *gin.Context) {
	_, ok := getContext[oauth2.TokenInfo](c, accessInfo)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewErrorsWithMessage(http.StatusText(http.StatusUnauthorized)))
		return
	}
	req := request.RequestPathForGenerateQRCode{
		Key: c.Param("key"),
	}
	if !validation(c, &req) {
		return
	}
	ctx := c.Request.Context()
	qrCode, err := h.service.GenerateQRCode(ctx, req.Key)
	if err != nil {
		logging.Default().Error(err)
		var clientErr *service.ClientError
		if errors.As(err, &clientErr) {
			errRes := response.NewErrorsWithMessage(clientErr.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, &errRes)
			return
		}
		errRes := response.NewErrorsWithMessage("An error occurred during URL generation.")
		c.AbortWithStatusJSON(http.StatusInternalServerError, errRes)
		return
	}
	writer := new(bytes.Buffer)
	teeReader := io.TeeReader(qrCode, writer)
	n, err := io.Copy(io.Discard, teeReader)
	if err != nil {
		logging.Default().Error(err)
		errRes := response.NewErrorsWithMessage("QR Code generation failed.")
		c.AbortWithStatusJSON(http.StatusInternalServerError, errRes)
		return
	}
	c.DataFromReader(http.StatusOK, n, "image/png", writer, nil)
}

// Remove 短縮URLの削除
//
// @Summary 短縮URLの削除
// @Tags API
// @Produce json
// @Param key path string true "short url key"
// @Success 204
// @Failure 400 {object} response.Errors
// @Failure 401 {object} response.Errors
// @Failure 500 {object} response.Errors
// @Router /shorts/{key} [delete]
// @Security OAuth2Application
// @ID RemoveShortURL
func (h *ShortURLHandler) Remove(c *gin.Context) {
	token, ok := getContext[oauth2.TokenInfo](c, accessInfo)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewErrorsWithMessage(http.StatusText(http.StatusUnauthorized)))
		return
	}
	key := h.getKey(c)
	v, err := key.Value()
	if err != nil {
		errRes := response.NewErrorsWithMessage("Invalid request")
		c.AbortWithStatusJSON(http.StatusBadRequest, errRes)
		return
	}
	ctx := c.Request.Context()
	err = h.service.Remove(ctx, v, token.GetUserID())
	if err != nil {
		if !errors.Is(err, service.ErrNotFound) {
			logging.Default().Error(err)
			errRes := response.NewErrorsWithMessage("An error occurred while deleting the URL.")
			c.JSON(http.StatusInternalServerError, errRes)
			return
		}
	}
	// 存在しないコードで削除を行われていても4xxではなく正常とみなす
	c.Status(http.StatusNoContent)
}

// Shorts 短縮URLの一覧
//
// @Summary 短縮URLの一覧
// @Tags API
// @Produce json
// @Success 200 {object} response.Shorts
// @Failure 400 {object} response.Errors
// @Failure 401 {object} response.Errors
// @Failure 500 {object} response.Errors
// @Router /shorts [get]
// @Security OAuth2Application
// @ID ShortURLList
func (h *ShortURLHandler) Shorts(c *gin.Context) {
	token, ok := getContext[oauth2.TokenInfo](c, accessInfo)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewErrorsWithMessage(http.StatusText(http.StatusUnauthorized)))
		return
	}
	ctx := c.Request.Context()
	values, err := h.service.FindAll(ctx, token.GetUserID())
	if err != nil {
		logging.Default().Error(err)
		errRes := response.NewErrorsWithMessage("An error occurred while deleting the URL.")
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	results := make([]response.Shorts, len(values))
	for idx, value := range values {
		v := response.Shorts{
			Key:       value.GetKey(),
			URL:       value.GetURL(),
			Author:    value.GetAuthor(),
			CreatedAt: utils.TimeToString(value.CreatedAt),
			UpdatedAt: utils.TimeToString(value.UpdatedAt),
		}
		results[idx] = v
	}
	c.JSON(http.StatusOK, &results)
}

func (h *ShortURLHandler) APIRouter(route gin.IRouter, middleware ...gin.HandlerFunc) {
	g := route.Group("/shorts")
	{
		g.Use(middleware...)
		g.POST("/generate", h.GenerateShortURL)

		g.DELETE("/:key", h.Remove)
		g.GET("/:key/qrcode", h.GenerateQRCode)
	}
}

func (h *ShortURLHandler) ServiceRouter(route gin.IRouter, middleware ...gin.HandlerFunc) {
	route.GET("/:key", h.Redirect)
}

func (h *ShortURLHandler) redirect(c *gin.Context, err error) {
	logging.Default().Error(err)
	c.Redirect(http.StatusTemporaryRedirect, utils.MustURL(h.option.appConfig.BaseURL, "/notfound"))
}

func (h *ShortURLHandler) getKey(c *gin.Context) types.String {
	return types.String(c.Param("key"))
}
