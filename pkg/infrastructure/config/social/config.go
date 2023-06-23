package social

import (
	"context"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/n-creativesystem/short-url/pkg/domain/social"
	"golang.org/x/oauth2"
)

type OAuthConfig struct {
	config    *oauth2.Config
	provider  *oidc.Provider
	claimKeys social.ClaimKeys
}

func (cfg *OAuthConfig) GetConfig() *social.Config {
	if cfg.config == nil {
		return nil
	}
	return &social.Config{
		Oauth2Config: cfg.config,
		Provider:     cfg.provider,
		ClaimKeys:    cfg.claimKeys,
	}
}

type socialConfig struct {
	ClientId     string `env:"CLIENT_ID"`
	ClientSecret string `env:"CLIENT_SECRET"`
	RedirectURL  string `env:"REDIRECT_URL"`
	Issuer       string `env:"ISSUER"`
	Username     string `env:"USERNAME"`
	Picture      string `env:"PICTURE"`
}

func NewConfig(ctx context.Context, prefix string) (*OAuthConfig, error) {
	cfg := socialConfig{}
	err := envProcess(&cfg, prefix)
	if err != nil {
		return nil, err
	}
	provider, err := oidc.NewProvider(ctx, cfg.Issuer)
	if err != nil {
		return nil, err
	}
	return &OAuthConfig{
		config: &oauth2.Config{
			ClientID:     cfg.ClientId,
			ClientSecret: cfg.ClientSecret,
			RedirectURL:  cfg.RedirectURL,
			Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
			Endpoint:     provider.Endpoint(),
		},
		provider: provider,
		claimKeys: social.ClaimKeys{
			Username: cfg.Username,
			Picture:  cfg.Picture,
		},
	}, nil
}
