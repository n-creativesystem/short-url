package sentry

import (
	"log/slog"
	"os"
	"sync"

	"github.com/getsentry/sentry-go"
	"github.com/n-creativesystem/short-url/pkg/utils"
	"github.com/n-creativesystem/short-url/pkg/utils/apps"
	"github.com/n-creativesystem/short-url/pkg/utils/logging"
)

var (
	sentryInit sync.Once
)

func IsEnable() bool {
	return utils.GetBoolEnv("SENTRY_ENABLED")
}

func Init() {
	option := sentry.ClientOptions{}
	sentryInit.Do(func() {
		option.Dsn = os.Getenv("SENTRY_DSN")
		option.Environment = apps.TrackingEnvironment()
		if err := sentry.Init(option); err != nil {
			slog.With(logging.WithErr(err)).Error("Sentry initialize")
		}
	})
}
