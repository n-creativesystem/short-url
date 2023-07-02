package handler

import (
	"os"

	"github.com/n-creativesystem/short-url/pkg/domain/config"
)

type ShortOption interface {
	apply(*shortOption)
}

type shortOptionFn func(*shortOption)

func (fn shortOptionFn) apply(option *shortOption) {
	fn(option)
}

func WithAppConfig(cfg *config.Application) ShortOption {
	return shortOptionFn(func(so *shortOption) {
		so.appConfig = cfg
	})
}

func newShortOption() shortOption {
	return shortOption{
		appConfig: &config.Application{
			BaseURL: os.Getenv("SERVICE_URL"),
		},
	}
}

type shortOption struct {
	appConfig *config.Application
}
