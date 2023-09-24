package handler

import (
	"io"
	"log/slog"
)

type Handle interface {
	slog.Handler
	io.Closer
}
