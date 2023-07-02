package logging

import (
	"context"
	"sync"

	"github.com/n-creativesystem/short-url/pkg/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

var (
	logger        *zap.Logger
	defaultLogger *zap.SugaredLogger
	once          sync.Once
)

func Logger() *zap.Logger {
	return logger
}

func sugaredLogger() *zap.SugaredLogger {
	if defaultLogger == nil {
		once.Do(func() {
			if logger == nil {
				SetFormat("console")
			}
			defaultLogger = logger.Sugar()
		})
	}
	return defaultLogger
}

func Default() *SentryLogger {
	return &SentryLogger{SugaredLogger: sugaredLogger().WithOptions(zap.AddCallerSkip(1))}
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

type contextKey int

const ctxKey contextKey = iota

func SetContext(ctx context.Context) context.Context {
	var log *SentryLogger
	span, ok := tracer.SpanFromContext(ctx)
	if ok {
		// NOTE: https://github.com/DataDog/dd-trace-go/blob/fbd37ea37336e8122aeccb6213b070041c4061e5/contrib/sirupsen/logrus/logrus.go#L31-L39
		spanContext := span.Context()
		log = &SentryLogger{SugaredLogger: sugaredLogger().With(
			zap.Uint64("dd.trace_id", spanContext.TraceID()),
			zap.Uint64("dd.span_id", spanContext.SpanID()),
		)}
	} else {
		log = &SentryLogger{SugaredLogger: sugaredLogger()}
	}
	return context.WithValue(ctx, ctxKey, log)
}

func Context(ctx context.Context) *SentryLogger {
	v, ok := ctx.Value(ctxKey).(*SentryLogger)
	if ok {
		return v
	}
	return Default()
}
