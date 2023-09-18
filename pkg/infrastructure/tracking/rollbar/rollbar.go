package rollbar

import (
	"os"
	"strconv"

	"github.com/rollbar/rollbar-go"
)

func IsEnable() bool {
	e := os.Getenv("ROLLBAR_ENABLED")
	b, err := strconv.ParseBool(e)
	return err == nil && b
}

func ErrorTracking(opts ...Option) *rollbar.Client {
	o := defaultOption()
	for _, opt := range opts {
		opt.apply(&o)
	}
	return rollbar.NewAsync(o.token, o.environment, o.codeVersion, "", o.serverRoot)
}
