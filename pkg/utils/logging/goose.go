package logging

import (
	"github.com/pressly/goose/v3"
)

func NewGooseLogger(l Log) *GooseLogger {
	return &GooseLogger{Log: l}
}

type GooseLogger struct {
	Log
}

var (
	_ goose.Logger = (*GooseLogger)(nil)
)

func (l *GooseLogger) Print(v ...interface{}) {
	l.Log.Info(v...)
}

func (l *GooseLogger) Printf(format string, v ...interface{}) {
	l.Log.Infof(format, v...)
}

func (l *GooseLogger) Println(v ...interface{}) {
	l.Log.Infoln(v...)
}
