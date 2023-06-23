package social

import (
	"bytes"
	"encoding/json"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/n-creativesystem/short-url/pkg/utils/logging"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

type ClaimKeys struct {
	Username string
	Picture  string
}

type Config struct {
	Oauth2Config *oauth2.Config
	Provider     *oidc.Provider
	ClaimKeys    ClaimKeys
}

type User struct {
	*oidc.UserInfo
	Username string `json:"username"`
	Picture  string `json:"picture"`
}

func (user *User) Encode() string {
	buf := new(bytes.Buffer)
	_ = json.NewEncoder(buf).Encode(&user)
	return buf.String()
}

func Decode(value string) (*User, error) {
	var user User
	buf := bytes.NewReader([]byte(value))
	if err := json.NewDecoder(buf).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (user *User) ParseClaims(cfg ClaimKeys) error {
	mapClaim := make(map[string]interface{})
	if err := user.UserInfo.Claims(&mapClaim); err != nil {
		logging.Default().Warn(err)
		return err
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(&mapClaim); err != nil {
		logging.Default().Warn(err)
		return err
	}
	//note: キーがネストしててもパースして取得できるので viper 使って楽した
	v := viper.New()
	v.SetConfigType("json")
	if err := v.ReadConfig(buf); err != nil {
		logging.Default().Warn(err)
		return err
	}
	if cfg.Username == "" {
		user.Username = user.Email
	} else {
		if v := v.GetString(cfg.Username); v != "" {
			user.Username = v
		}
	}
	if cfg.Picture != "" {
		if v := v.GetString(cfg.Picture); v != "" {
			user.Picture = v
		}
	}
	return nil
}
