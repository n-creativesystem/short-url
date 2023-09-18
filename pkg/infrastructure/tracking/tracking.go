package tracking

import (
	"context"
	"log/slog"
	"os"

	"github.com/n-creativesystem/short-url/pkg/infrastructure/tracking/rollbar"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/tracking/sentry"
	"github.com/n-creativesystem/short-url/pkg/utils"
	"github.com/n-creativesystem/short-url/pkg/utils/logging"
	"github.com/n-creativesystem/short-url/pkg/utils/logging/handler"
)

func Init() func() {
	cleanUps := make([]func(), 0, 10)
	cleanUpFn := func(closer []func()) func() {
		return func() {
			for _, fn := range closer {
				fn()
			}
		}
	}
	loggingHandler := make([]slog.Handler, 0, 10)
	if utils.IsDevOrCIorTest() {
		loggingHandler = append(loggingHandler, handler.NewTextHandler())
	} else {
		loggingHandler = append(loggingHandler, handler.NewJSONHandler())
	}
	// setup error tracking
	if sentry.IsEnable() {
		sentry.Init()
		loggingHandler = append(loggingHandler, handler.NewSentryHandler())
	}
	if rollbar.IsEnable() {
		client := rollbar.ErrorTracking()
		cleanUps = append(cleanUps, func() { _ = client.Close() })
		loggingHandler = append(loggingHandler, handler.NewRollbarHandler(client))
	}
	// APM tracking
	_, cleanup, err := NewTracerProvider(context.Background(), os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT"))
	if err != nil {
		slog.With(logging.WithErr(err)).Error("OTEL setup.")
	} else {
		cleanUps = append(cleanUps, cleanup)
	}

	logging.NewLogger(loggingHandler...)
	return cleanUpFn(cleanUps)
}
