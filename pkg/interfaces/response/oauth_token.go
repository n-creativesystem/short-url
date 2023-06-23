package response

type OAuth2Token struct {
	AccessToken  string `json:"access_token,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
	ExpiresIn    int64  `json:"expires_in,omitempty"`
	Scope        string `json:"scope,omitempty"`         // nullable
	RefreshToken string `json:"refresh_token,omitempty"` // nullable
}
