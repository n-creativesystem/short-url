package handler

import (
	"context"
	"log/slog"
	"slices"

	"go.uber.org/multierr"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type datadogHandler struct {
	handlers []slog.Handler
}

func NewDatadogHandler(handlers ...slog.Handler) slog.Handler {
	return &datadogHandler{
		handlers: handlers,
	}
}

func (h *datadogHandler) handler(fn func(h slog.Handler)) {
	for _, handler := range h.handlers {
		fn(handler)
	}
}

func (h *datadogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	flags := make([]bool, 0, len(h.handlers))
	h.handler(func(h slog.Handler) {
		flags = append(flags, h.Enabled(ctx, level))
	})
	return slices.Contains(flags, true)
}

func (h *datadogHandler) Handle(ctx context.Context, record slog.Record) error {
	span, ok := tracer.SpanFromContext(ctx)
	if ok {
		spanContext := span.Context()
		record.AddAttrs(slog.Uint64("dd.trace_id", spanContext.TraceID()))
		record.AddAttrs(slog.Uint64("dd.span_id", spanContext.SpanID()))
	}
	var err error
	h.handler(func(h slog.Handler) {
		if h.Enabled(ctx, record.Level) {
			if e := h.Handle(ctx, record); e != nil {
				err = multierr.Append(err, e)
			}
		}
	})
	return err
}

func (h *datadogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	handlers := make([]slog.Handler, 0, len(h.handlers))
	h.handler(func(h slog.Handler) {
		handlers = append(handlers, h.WithAttrs(attrs))
	})
	return NewDatadogHandler(handlers...)
}

func (h *datadogHandler) WithGroup(name string) slog.Handler {
	handlers := make([]slog.Handler, 0, len(h.handlers))
	h.handler(func(h slog.Handler) {
		handlers = append(handlers, h.WithGroup(name))
	})
	return NewDatadogHandler(handlers...)
}
