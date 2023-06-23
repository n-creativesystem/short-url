package social

import (
	"context"
)

type Repository interface {
	GetProvider(ctx context.Context, providerName string) (*Config, error)
	GetProviders(ctx context.Context) (map[string]*Config, error)
}
