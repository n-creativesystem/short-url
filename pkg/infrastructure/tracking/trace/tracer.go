package trace

import (
	"context"
	"runtime"

	"go.opentelemetry.io/otel/trace"
)

type tracer struct {
	trace.Tracer
}

func Tracer(t trace.Tracer) trace.Tracer {
	return &tracer{Tracer: t}
}

func (t *tracer) Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	if spanName == "" {
		spanName = getFuncName(1)
	}
	return t.Tracer.Start(ctx, spanName, opts...)
}

func getFuncName(skip int) string {
	pt, _, _, ok := runtime.Caller(skip + 1)
	if !ok {
		return "unknown func"
	}
	return runtime.FuncForPC(pt).Name()
}
