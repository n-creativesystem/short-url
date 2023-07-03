package social

import (
	"testing"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/stretchr/testify/require"
)

func TestUnsafeAddr(t *testing.T) {
	u := User{
		UserInfo: &oidc.UserInfo{},
	}
	u.SetClaims([]byte(`{"aaa":"test"}`))
	mp := map[string]string{}
	err := u.Claims(&mp)
	require.NoError(t, err)
	require.Equal(t, mp["aaa"], "test")
	require.NotEmpty(t, u.GetClaims())
}
