package rollbar

import (
	"os"

	"github.com/n-creativesystem/short-url/pkg/utils"
	"github.com/n-creativesystem/short-url/pkg/utils/apps"
	"github.com/rollbar/rollbar-go"
)

type option struct {
	token          string
	environment    string
	codeVersion    string
	serverRoot     string
	rollbarOptions []rollbar.OptionFunc
}

type Option interface {
	apply(opt *option)
}

type optionFn func(opt *option)

func defaultOption() option {
	return option{
		token:       os.Getenv("ROLLBAR_TOKEN"),
		environment: apps.TrackingEnvironment(),
		codeVersion: utils.Getenv("CODE_VERSION", "v1"),
		serverRoot:  apps.ServerRoot(),
	}
}

func (fn optionFn) apply(opt *option) {
	fn(opt)
}

func WithToken(token string) Option {
	return optionFn(func(opt *option) {
		opt.token = token
	})
}

func WithEnvironment(environment string) Option {
	return optionFn(func(opt *option) {
		opt.environment = environment
	})
}

func WithRollbarOption(rollbarOpt rollbar.OptionFunc) Option {
	return optionFn(func(opt *option) {
		opt.rollbarOptions = append(opt.rollbarOptions, rollbarOpt)
	})
}
