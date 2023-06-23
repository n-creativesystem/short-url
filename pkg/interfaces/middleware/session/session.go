package session

import (
	"net/http"
	"sync"
	"time"

	"github.com/alexedwards/scs/v2"
)

var (
	sessionManager *scs.SessionManager
	scsOnce        sync.Once
)

func initSessionManage(opts ...Option) {
	sessionManager = scs.New()
	for _, opt := range opts {
		opt.apply(sessionManager)
	}
}

func GetSessionManager(opts ...Option) *scs.SessionManager {
	defaultOpts := []Option{
		WithLifeTime((24 * time.Hour) * 10), // 10日間
		WithIdleTimeout(60 * time.Minute),   // 1時間
		WithCookieName("_short_url_session"),
		WithCookieHttpOnly(true),
		WithCookieDomain("localhost"),
		WithCookieSameSite(http.SameSiteLaxMode),
		WithCodec(JsonCodec{}),
		WithSecure(false),
	}
	defaultOpts = append(defaultOpts, opts...)
	scsOnce.Do(func() {
		initSessionManage(defaultOpts...)
	})
	return sessionManager
}
