//go:generate go run -mod=mod github.com/golang/mock/mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../mock/repository/$GOPACKAGE/$GOFILE
package social

import (
	"context"
)

type Repository interface {
	GetProvider(ctx context.Context, providerName string) (*Config, error)
	GetProviders(ctx context.Context) (map[string]*Config, error)
}

type UserRepository interface {
	Register(ctx context.Context, user *User) (*User, error)
	Login(ctx context.Context, email string) (*User, error)
}
