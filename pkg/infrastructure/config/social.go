package config

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	domain_social "github.com/n-creativesystem/short-url/pkg/domain/social"
	"github.com/n-creativesystem/short-url/pkg/utils/logging"
	"github.com/sethvargo/go-envconfig"

	. "github.com/n-creativesystem/short-url/pkg/infrastructure/config/social"
)

type socialEnvConfigImpl struct {
	Providers []string `env:"IDPs"`
}

func NewSocialConfig() domain_social.Repository {
	impl := &socialEnvConfigImpl{}
	_ = envconfig.Process(context.Background(), impl)
	return impl
}

var (
	_ domain_social.Repository = (*socialEnvConfigImpl)(nil)
)

func (cfg *socialEnvConfigImpl) GetProvider(ctx context.Context, providerName string) (*domain_social.Config, error) {
	p := strings.ToUpper(providerName)
	switch p {
	case "GOOGLE":
		issuer := "GOOGLE_ISSUER"
		if v := os.Getenv(issuer); v == "" {
			os.Setenv(issuer, "https://accounts.google.com")
		}
	}
	config, err := NewConfig(ctx, fmt.Sprintf("%s_", p))
	if err != nil {
		return nil, err
	}
	return config.GetConfig(), nil
}

func (cfg *socialEnvConfigImpl) GetProviders(ctx context.Context) (map[string]*domain_social.Config, error) {
	mp := make(map[string]*domain_social.Config, len(cfg.Providers))
	for _, provider := range cfg.Providers {
		if c, err := cfg.GetProvider(ctx, provider); err != nil {
			logging.Default().Warn(err)
		} else {
			mp[provider] = c
		}
	}
	if len(mp) == 0 {
		return nil, errors.New("Social login is not set.")
	}
	return mp, nil
}
