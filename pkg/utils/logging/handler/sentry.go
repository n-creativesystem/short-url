package handler

import (
	"log/slog"
	"os"
	"strings"

	slogsentry "github.com/samber/slog-sentry"
)

func getSentryLevel() slog.Level {
	level := os.Getenv("SENTRY_LOG_LEVEL")
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn", "warning":
		return slog.LevelWarn
	case "err", "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func NewSentryHandler() slog.Handler {
	option := slogsentry.Option{
		Level: getSentryLevel(),
	}
	return NewErrorTracking(NewAsyncHandler(option.NewSentryHandler()))
}
