package handler

import (
	"context"
	"log/slog"
)

type AsyncHandler struct {
	slog.Handler
}

func NewAsyncHandler(h slog.Handler) slog.Handler {
	return &AsyncHandler{Handler: h}
}

func (h *AsyncHandler) Handle(ctx context.Context, r slog.Record) error {
	go func() { _ = h.Handler.Handle(ctx, r) }()
	return nil
}

func (h *AsyncHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewAsyncHandler(h.Handler.WithAttrs(attrs))
}

func (h *AsyncHandler) WithGroup(name string) slog.Handler {
	return NewAsyncHandler(h.Handler.WithGroup(name))
}
