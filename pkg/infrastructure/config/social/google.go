package social

import (
	"context"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type googleConfig struct {
	Config socialConfig `env:",prefix=GOOGLE_"`
}

func NewGoogleConfig() (*OAuthConfig, error) {
	cfg := googleConfig{}
	err := envProcess(&cfg, "GOOGLE_")
	if err != nil {
		return nil, err
	}
	provider, err := oidc.NewProvider(context.Background(), "https://accounts.google.com")
	if err != nil {
		return nil, err
	}
	return &OAuthConfig{
		config: &oauth2.Config{
			ClientID:     cfg.Config.ClientId,
			ClientSecret: cfg.Config.ClientSecret,
			RedirectURL:  cfg.Config.RedirectURL,
			Scopes:       []string{oidc.ScopeOpenID, "email"},
			Endpoint:     google.Endpoint,
		},
		provider: provider,
	}, nil
}
