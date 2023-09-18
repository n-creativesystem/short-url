package middleware

import (
	"bufio"
	"bytes"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/gin-gonic/gin"
	"github.com/n-creativesystem/short-url/pkg/domain/social"
	. "github.com/n-creativesystem/short-url/pkg/interfaces/middleware/session"
	"github.com/n-creativesystem/short-url/pkg/interfaces/response"
	"github.com/n-creativesystem/short-url/pkg/utils/logging"
)

var (
	ErrAuthorize = errors.New("Authorize")
)

func Session(opts ...Option) gin.HandlerFunc {
	s := GetSessionManager(opts...)
	return func(c *gin.Context) {
		w := c.Writer
		r := c.Request
		var token string
		cookie, err := r.Cookie(s.Cookie.Name)
		if err == nil {
			token = cookie.Value
		}

		ctx, err := s.Load(r.Context(), token)
		if err != nil {
			msg := "Session check failed."
			slog.With(logging.WithErr(err)).ErrorContext(ctx, msg)
			c.AbortWithStatusJSON(http.StatusInternalServerError, response.NewErrorsWithMessage(http.StatusText(http.StatusInternalServerError)))
			return
		}
		SetGinContext(c)
		ctx = SetContext(ctx)
		*c.Request = *r.WithContext(ctx)
		bw := &bufferedResponseWriter{ResponseWriter: w}
		c.Writer = bw
		c.Next()

		if c.Request.MultipartForm != nil {
			_ = c.Request.MultipartForm.RemoveAll()
		}

		switch s.Status(ctx) {
		case scs.Modified:
			_, _, err := s.Commit(ctx)
			if err != nil {
				msg := "Session check failed."
				slog.With(logging.WithErr(err)).ErrorContext(ctx, msg)
				c.AbortWithStatusJSON(http.StatusInternalServerError, response.NewErrorsWithMessage(http.StatusText(http.StatusInternalServerError)))
				return
			}
			token, expiry, err := s.Commit(ctx)
			if err != nil {
				msg := "Session check failed."
				slog.With(logging.WithErr(err)).ErrorContext(ctx, msg)
				c.AbortWithStatusJSON(http.StatusInternalServerError, response.NewErrorsWithMessage(http.StatusText(http.StatusInternalServerError)))
				return
			}
			s.WriteSessionCookie(ctx, bw, token, expiry)
		case scs.Destroyed:
			s.WriteSessionCookie(ctx, w, "", time.Time{})
		}

		w.Header().Add("Vary", "Cookie")

		if bw.code != 0 {
			w.WriteHeader(bw.code)
		}
		_, _ = w.Write(bw.buf.Bytes())
	}
}

func UnauthorizeRedirect(redirect string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				if errors.Is(err, ErrAuthorize) {
					c.Redirect(http.StatusFound, redirect)
					return
				}
			}
		}
	}
}

func Protected(repo social.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		sm := GetContext(ctx)
		buf := sm.GetString(ctx, LoginUser)
		user, err := social.Decode(buf)
		if err != nil {
			_ = sm.Destroy(ctx)
			_ = c.Error(ErrAuthorize)
			c.Abort()
			return
		}
		user, err = repo.Login(ctx, user.Email)
		if err != nil {
			_ = sm.Destroy(ctx)
			_ = c.Error(ErrAuthorize)
			c.Abort()
			return
		}
		SetAuthUser(c, user)
		c.Next()
	}
}

type bufferedResponseWriter struct {
	gin.ResponseWriter
	buf         bytes.Buffer
	code        int
	wroteHeader bool
}

func (bw *bufferedResponseWriter) Write(b []byte) (int, error) {
	return bw.buf.Write(b)
}

func (bw *bufferedResponseWriter) WriteHeader(code int) {
	if !bw.wroteHeader && code >= 100 {
		bw.code = code
		bw.wroteHeader = true
	}
}

func (bw *bufferedResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hj := bw.ResponseWriter.(http.Hijacker)
	return hj.Hijack()
}

func (bw *bufferedResponseWriter) Push(target string, opts *http.PushOptions) error {
	if pusher, ok := bw.ResponseWriter.(http.Pusher); ok {
		return pusher.Push(target, opts)
	}
	return http.ErrNotSupported
}

func (bw *bufferedResponseWriter) WriteHeaderNow() {
	slog.Debug("WriteHeaderNow")
}
