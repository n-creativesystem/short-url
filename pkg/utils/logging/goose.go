package logging

import (
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
)

func NewGooseLogger(l *zap.SugaredLogger) *GooseLogger {
	return &GooseLogger{SugaredLogger: l}
}

type GooseLogger struct {
	*zap.SugaredLogger
}

var (
	_ goose.Logger = (*GooseLogger)(nil)
)

func (l *GooseLogger) Print(v ...interface{}) {
	l.SugaredLogger.Info(v...)
}

func (l *GooseLogger) Printf(format string, v ...interface{}) {
	l.SugaredLogger.Infof(format, v...)
}

func (l *GooseLogger) Println(v ...interface{}) {
	l.SugaredLogger.Infoln(v...)
}
