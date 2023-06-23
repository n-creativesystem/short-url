package response

type RegisterOAuth2Client struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type OAuth2Application struct {
	ClientID string `json:"client_id"`
	Name     string `json:"name"`
}

type OAuth2Applications struct {
	Results []OAuth2Application `json:"results"`
}
