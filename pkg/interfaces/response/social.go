package response

import "github.com/coreos/go-oidc/v3/oidc"

type EnabledSocialLogin struct {
	Socials []string `json:"socials"`
}

type User struct {
	*oidc.UserInfo
	Username string `json:"username"`
	Picture  string `json:"picture"`
}
