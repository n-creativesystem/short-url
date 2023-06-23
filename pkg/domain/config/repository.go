//go:generate go run -mod=mod github.com/golang/mock/mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../mock/repository/$GOPACKAGE/$GOFILE
package config

import "context"

type ApplicationRepository interface {
	Get(ctx context.Context, opts ...OptionFunc) (*Application, error)
}
