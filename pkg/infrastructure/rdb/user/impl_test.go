package user

import (
	"context"
	"testing"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/n-creativesystem/short-url/pkg/domain/config"
	"github.com/n-creativesystem/short-url/pkg/domain/social"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/tests"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials/crypto"
	"github.com/stretchr/testify/require"
)

func TestRegister(t *testing.T) {
	err := crypto.Init()
	require.NoError(t, err)
	info := &oidc.UserInfo{
		Subject:       "subject",
		Profile:       "profile",
		Email:         "test@example.com",
		EmailVerified: false,
	}
	value := social.User{
		UserInfo: info,
		Username: "username",
		Picture:  "picture",
	}
	value.SetClaims([]byte(`{"aaa": "test"}`))
	require := require.New(t)
	t.Parallel()
	for _, d := range config.RDB {
		d := d
		t.Run(d.String(), func(t *testing.T) {
			config.SetDriver(d)
			client := tests.GetTestDBAndMigrate("shorturl_by_short")
			rdb.SetClient(client)
			ctx := context.Background()
			repoImpl := newRepository()
			user, err := repoImpl.Register(ctx, &value)
			require.NoError(err)
			findUser, err := repoImpl.Login(ctx, value.Email)
			require.NoError(err)
			require.EqualValues(value, *user)
			require.EqualValues(value, *findUser)
			info := *value.UserInfo
			modifyValue := value
			modifyValue.UserInfo = &info
			modifyValue.Subject = "update_subject"
			modifyValue.Profile = "update_profile"
			modifyValue.EmailVerified = true
			modifyValue.Username = "update_username"
			modifyValue.Picture = "update_picture"
			user, err = repoImpl.Register(ctx, &modifyValue)
			require.NoError(err)
			findUser, err = repoImpl.Login(ctx, value.Email)
			require.NoError(err)
			require.NotEqualValues(value, *user)
			require.EqualValues(modifyValue, *user)
			require.NotEqualValues(value, *findUser)
			require.EqualValues(modifyValue, *findUser)
		})
	}
}
