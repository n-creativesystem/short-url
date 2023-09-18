package social

import (
	"bytes"
	"encoding/json"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/n-creativesystem/short-url/pkg/utils"
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
		return err
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(&mapClaim); err != nil {
		return err
	}
	//note: キーがネストしててもパースして取得できるので viper 使って楽した
	v := viper.New()
	v.SetConfigType("json")
	if err := v.ReadConfig(buf); err != nil {
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

func (u *User) EqualSubject(value string) bool {
	return u.Subject == value
}

func (u *User) EqualProfile(value string) bool {
	return u.Profile == value
}

func (u *User) EqualEmailVerified(value bool) bool {
	return u.EmailVerified == value
}

func (u *User) EqualUsername(value string) bool {
	return u.Username == value
}

func (u *User) EqualPicture(value string) bool {
	return u.Picture == value
}

func (u *User) SetClaims(data []byte) {
	updateClaims(u.UserInfo, data)
}

func (u *User) GetClaims() []byte {
	return getClaims(u.UserInfo)
}

func updateClaims(value *oidc.UserInfo, data []byte) {
	utils.UpdateField(value, "claims", data)
}

func getClaims(value *oidc.UserInfo) []byte {
	return utils.GetField[[]byte](value)
}
