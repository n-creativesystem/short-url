package middleware

import (
	"errors"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	pkgErr "github.com/n-creativesystem/short-url/pkg/utils/errors"
	"github.com/n-creativesystem/short-url/pkg/utils/logging/handler"
)

func Logger(notLogged ...string) gin.HandlerFunc {
	timeFormat := time.RFC3339
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknow"
	}

	var skip map[string]struct{}

	if length := len(notLogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, p := range notLogged {
			skip[p] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx, span := tracer.Start(ctx, "LoggerMiddleware")
		defer span.End()
		*c.Request = *c.Request.WithContext(ctx)
		path := c.Request.URL.Path
		if _, ok := skip[path]; ok {
			return
		}
		start := time.Now()
		c.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}
		entry := slog.With(
			"hostname", hostname,
			"statusCode", statusCode,
			"latency", fmt.Sprintf("%dms", latency),
			"clientIP", clientIP,
			"method", c.Request.Method,
			"path", path,
			"referer", referer,
			"dataLength", dataLength,
			"userAgent", clientUserAgent,
			"time", time.Now().Format(timeFormat),
		)
		if len(c.Errors) > 0 {
			if IsIgnoreError(c.Errors) {
				entry = entry.With(handler.IgnoreTracing)
			}
			entry.With("err", c.Errors).ErrorContext(ctx, c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := "Request"
			if statusCode >= http.StatusInternalServerError {
				entry.ErrorContext(ctx, msg)
			} else if statusCode >= http.StatusBadRequest {
				entry.WarnContext(ctx, msg)
			} else {
				entry.InfoContext(ctx, msg)
			}
		}
	}
}

func IsIgnoreError(ginErrs []*gin.Error) bool {
	var ignoreErr *pkgErr.IgnoreError
	for _, err := range ginErrs {
		if errors.As(err, &ignoreErr) {
			return true
		}
	}
	return false
}
