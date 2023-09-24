package logging

import (
	"context"
	"log/slog"
)

var (
	logKey struct{}
)

func ToContext(ctx context.Context, log *slog.Logger) context.Context {
	if log == nil {
		log = FromContext(ctx)
	}
	return context.WithValue(ctx, logKey, log)
}

func FromContext(ctx context.Context) *slog.Logger {
	if v, ok := ctx.Value(logKey).(*slog.Logger); ok {
		return v
	}
	return slog.Default()
}
