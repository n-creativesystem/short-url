package request

type RevokeToken struct {
	Token         string `json:"token"`
	TokenTypeHint string `json:"token_type_hint,omitempty"` // Enums(access_token,refresh_token)
}
