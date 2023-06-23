package session

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n-creativesystem/short-url/pkg/domain/social"
)

const LoginUser = "loginUser"

type loginUserContextKey struct{}

var loginUserKey loginUserContextKey

func AuthLogin(ctx context.Context, user *social.User) {
	sm := GetContext(ctx)
	sm.Put(ctx, LoginUser, user.Encode())
}

func AuthLogout(ctx context.Context) {
	sm := GetContext(ctx)
	sm.Remove(ctx, LoginUser)
}

func SetAuthUser(c *gin.Context, user *social.User) {
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, loginUserKey, user)
	*c.Request = *c.Request.WithContext(ctx)
	c.Set(LoginUser, user)
}

func GetAuthUserWithContext(ctx context.Context) (*social.User, bool) {
	v, ok := ctx.Value(loginUserKey).(*social.User)
	return v, ok
}

func GetAuthUserWithGinContext(c *gin.Context) (*social.User, bool) {
	if value, ok := c.Get(LoginUser); ok {
		v, ok := value.(*social.User)
		if ok {
			return v, true
		}
	}
	c.AbortWithStatus(http.StatusUnauthorized)
	return nil, false
}
