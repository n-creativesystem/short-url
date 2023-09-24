package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n-creativesystem/short-url/pkg/interfaces/handler/csrf"
	"github.com/n-creativesystem/short-url/pkg/interfaces/response"
)

const (
	csrf_name = "csrf_token"
)

type TokenGetter func(c *gin.Context) string

var (
	defaultRequestTokenGetter = func(c *gin.Context) string {
		r := c.Request
		if t := r.FormValue("_csrf"); t != "" {
			return t
		}
		if t := r.URL.Query().Get("_csrf"); t != "" {
			return t
		}
		if t := r.Header.Get("X-CSRF-TOKEN"); t != "" {
			return t
		}
		if t := r.Header.Get("X-XSRF-TOKEN"); t != "" {
			return t
		}
		return ""
	}

	defaultSessionTokenGetter = func(c *gin.Context) string {
		r := c.Request
		if c, err := r.Cookie(csrf_name); err == nil {
			return c.Value
		}
		return ""
	}
)

type CSRFTokenHandler struct {
	ignoreMethod       []string
	tokenGetter        TokenGetter
	sessionTokenGetter TokenGetter
}

func NewCSRFTokenHandler(opts ...CSRFTokenOption) *CSRFTokenHandler {
	ignoreMethods := []string{http.MethodGet, http.MethodOptions, http.MethodHead}
	h := &CSRFTokenHandler{
		ignoreMethod:       ignoreMethods,
		tokenGetter:        defaultRequestTokenGetter,
		sessionTokenGetter: defaultSessionTokenGetter,
	}
	for _, opt := range opts {
		opt.apply(h)
	}
	return h
}

// GetToken
//
// @Summary CSRFトークンの生成
// @Tags UI
// @Accept json
// @Produce json
// @Success 200 {object} response.CsrfToken
// @Router /csrf_token [get]
// @ID GetCSRFToken
func (h *CSRFTokenHandler) GetToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx, span := tracer.Start(ctx, "")
		defer span.End()
		*c.Request = *c.Request.WithContext(ctx)
		var resp response.CsrfToken
		token, maskToken := csrf.GenerateToken()
		cookie := &http.Cookie{
			Name:     csrf_name,
			Value:    maskToken,
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			Secure:   true,
			MaxAge:   3600 * 12,
		}
		http.SetCookie(c.Writer, cookie)
		resp.CsrfToken = token
		c.JSON(http.StatusOK, &resp)
	}
}

func (h *CSRFTokenHandler) Middleware() gin.HandlerFunc {
	mpIgnoreMethod := make(map[string]struct{}, len(h.ignoreMethod))
	for _, v := range h.ignoreMethod {
		mpIgnoreMethod[v] = struct{}{}
	}
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx, span := tracer.Start(ctx, "")
		defer span.End()
		*c.Request = *c.Request.WithContext(ctx)
		r := c.Request
		if _, ok := mpIgnoreMethod[r.Method]; ok {
			c.Next()
			return
		}
		token := h.tokenGetter(c)
		sessionToken := h.sessionTokenGetter(c)

		if !csrf.VerifyToken(token, sessionToken) {
			c.AbortWithStatusJSON(http.StatusForbidden, response.NewErrorsWithMessage("Invalid csrf token"))
			return
		}
		c.Next()
	}
}

type (
	CSRFTokenOption interface {
		apply(*CSRFTokenHandler)
	}
	csrfTokenFn func(*CSRFTokenHandler)
)

func (fn csrfTokenFn) apply(h *CSRFTokenHandler) {
	fn(h)
}

func WithIgnoreMethods(methods ...string) CSRFTokenOption {
	return csrfTokenFn(func(ch *CSRFTokenHandler) {
		ch.ignoreMethod = methods
	})
}

func WithTokenGetter(getter TokenGetter) CSRFTokenOption {
	return csrfTokenFn(func(ch *CSRFTokenHandler) {
		ch.tokenGetter = getter
	})
}

func WithSessionTokenGetter(getter TokenGetter) CSRFTokenOption {
	return csrfTokenFn(func(ch *CSRFTokenHandler) {
		ch.sessionTokenGetter = getter
	})
}
