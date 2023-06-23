package logging

import (
	"sync"

	"github.com/n-creativesystem/short-url/pkg/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger        *zap.Logger
	defaultLogger *zap.SugaredLogger
	once          sync.Once
)

func Logger() *zap.Logger {
	return logger
}

func Default() *zap.SugaredLogger {
	if defaultLogger == nil {
		if logger == nil {
			SetFormat("console")
		}
		once.Do(func() {
			defaultLogger = logger.Sugar()
		})
	}
	return defaultLogger
}

func init() {
	if utils.IsTest() || utils.IsCI() {
		initialize(zap.NewDevelopmentConfig())
	}
}

func initialize(config zap.Config) {
	config.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "st",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	l, err := config.Build()
	if err != nil {
		panic(err)
	}
	logger = l
}

func SetFormat(format string) {
	if format == "console" {
		setLogger(zap.NewDevelopmentConfig())
	} else {
		setLogger(zap.NewProductionConfig())
	}
}

func DebugMode() {
	setLogger(zap.NewDevelopmentConfig())
}

type Options func(*zap.Config)

func setLogger(config zap.Config, opts ...Options) {
	Close()
	for _, opt := range opts {
		opt(&config)
	}
	initialize(config)
}

func Close() {
	if logger := Logger(); logger != nil {
		_ = logger.Sync()
	}
}

func ParseLogLevel(level string) (zap.AtomicLevel, error) {
	return zap.ParseAtomicLevel(level)
}
