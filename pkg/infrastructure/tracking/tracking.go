package tracking

import (
	"log/slog"
	"os"
	"strings"

	"github.com/n-creativesystem/short-url/pkg/infrastructure/tracking/datadog"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/tracking/rollbar"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/tracking/sentry"
	"github.com/n-creativesystem/short-url/pkg/utils"
	"github.com/n-creativesystem/short-url/pkg/utils/logging"
	"github.com/n-creativesystem/short-url/pkg/utils/logging/handler"
)

func Init() func() {
	closer := make([]func(), 0, 10)
	closeFn := func(closer []func()) func() {
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
		closer = append(closer, func() { _ = client.Close() })
		loggingHandler = append(loggingHandler, handler.NewRollbarHandler(client))
	}
	// APM tracking
	switch strings.ToLower(os.Getenv("APM_TRACING")) {
	case "SENTRY":
		sentry.APM()
	case "DATADOG":
		closeFn := datadog.APM()
		closer = append(closer, closeFn)
		loggingHandler = append(loggingHandler, handler.NewDatadogHandler())
	}

	logging.NewLogger(loggingHandler...)
	return closeFn(closer)
}
