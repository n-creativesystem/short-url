package logging

import (
	"fmt"
	"log/slog"

	"github.com/pressly/goose/v3"
)

func NewGooseLogger() *GooseLogger {
	return &GooseLogger{Logger: slog.Default()}
}

type GooseLogger struct {
	*slog.Logger
}

var (
	_ goose.Logger = (*GooseLogger)(nil)
)

func (l *GooseLogger) Fatalf(format string, v ...interface{}) {
	l.Logger.Info(fmt.Sprintf(format, v...))
}

func (l *GooseLogger) Printf(format string, v ...interface{}) {
	l.Logger.Info(fmt.Sprintf(format, v...))
}
