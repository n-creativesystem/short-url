package oauth2

import (
	"time"

	oauth2client "github.com/n-creativesystem/short-url/pkg/domain/oauth2_client"
)

type tokenOption struct {
	ticker *time.Ticker
	client oauth2client.Repository
}

type TokenOption interface {
	apply(*tokenOption)
}

type tokenOptionFunc func(store *tokenOption)

func (f tokenOptionFunc) apply(store *tokenOption) {
	f(store)
}

func WithGCTimeInterval(interval int) TokenOption {
	return tokenOptionFunc(func(store *tokenOption) {
		if interval != 0 {
			store.ticker = time.NewTicker(time.Second * time.Duration(interval))
		}
	})
}

func WithOAuth2Client(repo oauth2client.Repository) TokenOption {
	return tokenOptionFunc(func(store *tokenOption) {
		store.client = repo
	})
}
