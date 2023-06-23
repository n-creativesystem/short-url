package session

import (
	"context"

	"github.com/alexedwards/scs/v2"
	"github.com/gin-gonic/gin"
)

type contextKey string

var ctxKey = contextKey("scs.sessionManger")

func SetGinContext(c *gin.Context) {
	c.Set(string(ctxKey), GetSessionManager())
}

func GetGinContext(c *gin.Context) *scs.SessionManager {
	v, ok := c.Get(string(ctxKey))
	if !ok {
		return GetSessionManager()
	}
	if value, ok := v.(*scs.SessionManager); ok {
		return value
	}
	return GetSessionManager()
}

func SetContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKey, GetSessionManager())
}

func GetContext(ctx context.Context) *scs.SessionManager {
	v, ok := ctx.Value(ctxKey).(*scs.SessionManager)
	if ok {
		return v
	}
	return GetSessionManager()
}
