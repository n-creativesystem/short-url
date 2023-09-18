package sentry

import (
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

func spanMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		span := trace.SpanFromContext(ctx)
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			spanContext := span.SpanContext()

			hub.Scope().SetTags(map[string]string{
				"trace_id": spanContext.TraceID().String(),
				"span_id":  spanContext.SpanID().String(),
			})
		}
		c.Next()
	}
}

func GinMiddleware(route *gin.Engine, opts ...GinOption) {
	o := defaultGinOption()
	for _, opt := range opts {
		opt.apply(&o)
	}
	route.Use(sentrygin.New(sentrygin.Options{
		Repanic:         o.rePanic,
		WaitForDelivery: o.waitForDelivery,
		Timeout:         o.timeout,
	}), spanMiddleware())
}
