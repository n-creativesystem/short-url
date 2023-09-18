package handler

import (
	"log/slog"
	"os"
	"strings"

	"github.com/rollbar/rollbar-go"
	slogrollbar "github.com/samber/slog-rollbar"
)

func getRollbarLevel() slog.Level {
	level := os.Getenv("ROLLBAR_LOG_LEVEL")
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

func NewRollbarHandler(client *rollbar.Client) slog.Handler {
	option := slogrollbar.Option{
		Level:  getRollbarLevel(),
		Client: client,
	}
	return option.NewRollbarHandler()
}
