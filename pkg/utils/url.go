package utils

import (
	"net/url"
)

func URL(baseURL string, paths ...string) (string, error) {
	return url.JoinPath(baseURL, paths...)
}

func MustURL(baseURL string, paths ...string) string {
	u, err := URL(baseURL, paths...)
	if err != nil {
		panic(err)
	}
	return u
}
