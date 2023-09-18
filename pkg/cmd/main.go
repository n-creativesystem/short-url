package cmd

import (
	"log/slog"
	"os"

	"github.com/n-creativesystem/short-url/pkg/infrastructure/tracking"
	"github.com/n-creativesystem/short-url/pkg/utils/logging"
)

func Execute() {
	tracingClose := tracking.Init()
	defer tracingClose()

	cmd := rootCommand()
	if err := cmd.Execute(); err != nil {
		slog.With(logging.WithErr(err)).Error(err.Error())
		os.Exit(1)
	}
}
