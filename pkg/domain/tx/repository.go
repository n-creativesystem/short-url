//go:generate go run -mod=mod github.com/golang/mock/mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../mock/repository/$GOPACKAGE/$GOFILE
package tx

import "context"

type Beginner func(opts ...OptionFunc) (ContextBeginner, error)

type ContextBeginner interface {
	BeginTx(ctx context.Context, fn func(ctx context.Context) error) error
}
