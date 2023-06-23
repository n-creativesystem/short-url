package handler

import (
	"os"

	"github.com/n-creativesystem/short-url/pkg/utils"
)

type ShortOption interface {
	apply(*shortOption)
}

type shortOptionFn func(*shortOption)

func (fn shortOptionFn) apply(option *shortOption) {
	fn(option)
}

func WithBaseURL(baseURL string) ShortOption {
	return shortOptionFn(func(so *shortOption) {
		so.baseURL = baseURL
	})
}

func newShortOption() shortOption {
	return shortOption{
		baseURL: os.Getenv("SERVICE_URL"),
	}
}

type shortOption struct {
	baseURL string
}

func (so *shortOption) URL(paths ...string) (string, error) {
	return utils.URL(so.baseURL, paths...)
}

func (so *shortOption) MustURL(paths ...string) string {
	return utils.MustURL(so.baseURL, paths...)
}
