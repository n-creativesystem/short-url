//go:generate go run -mod=mod github.com/golang/mock/mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../mock/repository/$GOPACKAGE/$GOFILE
package short

import (
	"context"
)

type Repository interface {
	Get(ctx context.Context, key string) (*Short, error)
	Put(ctx context.Context, value Short) error
	Del(ctx context.Context, key, author string) (bool, error)
	Exists(ctx context.Context, key string) (bool, error)
	FindAll(ctx context.Context, author string) ([]ShortWithTimeStamp, error)
	FindByKeyAndAuthor(ctx context.Context, key, author string) (*ShortWithTimeStamp, error)
}
