package logging

import (
	"log/slog"

	"github.com/n-creativesystem/short-url/pkg/utils/logging/handler"
)

func NewLogger(handlers ...slog.Handler) handler.Handle {
	return handler.NewHandler(handlers...)
}
