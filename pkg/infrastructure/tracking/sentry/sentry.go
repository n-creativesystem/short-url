package sentry

import (
	"log/slog"
	"os"
	"strings"
	"sync"

	"github.com/getsentry/sentry-go"
	sentryotel "github.com/getsentry/sentry-go/otel"
	"github.com/n-creativesystem/short-url/pkg/utils"
	"github.com/n-creativesystem/short-url/pkg/utils/logging"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var (
	sentryInit sync.Once
)

func IsEnable() bool {
	return utils.GetBoolEnv("SENTRY_ENABLED") || IsEnableAPM()
}

func IsEnableAPM() bool {
	return strings.EqualFold(os.Getenv("APM_TRACING"), "SENTRY")
}

func Init() {
	option := sentry.ClientOptions{}
	sentryInit.Do(func() {
		option.Dsn = os.Getenv("SENTRY_DSN")
		option.Environment = utils.Getenv("TRACKING_ENV", "local")
		if IsEnableAPM() {
			option.EnableTracing = IsEnableAPM()
			option.TracesSampleRate = 1.0
		}
		if err := sentry.Init(option); err != nil {
			slog.With(logging.WithErr(err)).Error("Sentry initialize")
		}
	})
}

func APM() {
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(sentryotel.NewSentrySpanProcessor()),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(sentryotel.NewSentryPropagator())
}
