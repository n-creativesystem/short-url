package rollbar

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/rollbar/rollbar-go"
	"go.opentelemetry.io/otel/trace"
)

func getCallers(skip int) (pc []uintptr) {
	pc = make([]uintptr, 1000)
	i := runtime.Callers(skip+1, pc)
	return pc[0:i]
}

func recoveryMiddleware(onlyCrashes bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rval := recover(); rval != nil {
				debug.PrintStack()

				rollbar.Critical(errors.New(fmt.Sprint(rval)), getCallers(3), map[string]string{
					"endpoint": c.Request.RequestURI,
				})

				c.AbortWithStatus(http.StatusInternalServerError)
			}

			if !onlyCrashes {
				var (
					traceId, spanId string
				)
				ctx := c.Request.Context()
				span := trace.SpanFromContext(ctx)
				if span != nil {
					spanContext := span.SpanContext()
					traceId = spanContext.TraceID().String()
					spanId = spanContext.SpanID().String()
				}
				for _, item := range c.Errors {
					mp := map[string]string{
						"meta":     fmt.Sprint(item.Meta),
						"endpoint": c.Request.RequestURI,
					}
					if traceId != "" {
						mp["dd.trace_id"] = traceId
					}
					if spanId != "" {
						mp["dd.span_id"] = spanId
					}
					rollbar.Error(item.Err, mp)
				}
			}
		}()

		c.Next()
	}
}

func GinMiddleware(route *gin.Engine, opts ...GinOption) {
	if !IsEnable() {
		return
	}
	o := defaultGinOption()
	for _, opt := range opts {
		opt.apply(&o)
	}
	route.Use(recoveryMiddleware(o.onlyCrashes))
}
