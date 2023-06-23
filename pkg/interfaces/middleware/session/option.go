package session

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/alexedwards/scs/v2/memstore"
)

type Option interface {
	apply(*scs.SessionManager)
}

type optionFn func(*scs.SessionManager)

func (fn optionFn) apply(sm *scs.SessionManager) {
	fn(sm)
}

func WithLifeTime(d time.Duration) Option {
	return optionFn(func(sm *scs.SessionManager) {
		sm.Lifetime = d
	})
}

func WithIdleTimeout(d time.Duration) Option {
	return optionFn(func(sm *scs.SessionManager) {
		sm.IdleTimeout = d
	})
}

func WithCookieDomain(domain string) Option {
	return optionFn(func(sm *scs.SessionManager) {
		sm.Cookie.Domain = domain
	})
}

func WithCookieName(name string) Option {
	return optionFn(func(sm *scs.SessionManager) {
		sm.Cookie.Name = name
	})
}

func WithCookieHttpOnly(value bool) Option {
	return optionFn(func(sm *scs.SessionManager) {
		sm.Cookie.HttpOnly = value
	})
}

func WithCookieSameSite(value http.SameSite) Option {
	return optionFn(func(sm *scs.SessionManager) {
		sm.Cookie.SameSite = value
	})
}

func WithCodec(codec scs.Codec) Option {
	if codec == nil {
		codec = JsonCodec{}
	}
	return optionFn(func(sm *scs.SessionManager) {
		sm.Codec = codec
	})
}

func WithSessionStore(store scs.Store) Option {
	if store == nil {
		store = memstore.New()
	}
	return optionFn(func(sm *scs.SessionManager) {
		sm.Store = store
	})
}

func WithSecure(value bool) Option {
	return optionFn(func(sm *scs.SessionManager) {
		sm.Cookie.Secure = value
	})
}
