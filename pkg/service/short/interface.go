//go:generate go run -mod=mod github.com/golang/mock/mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../mock/service/$GOPACKAGE/$GOFILE
package short

import (
	"context"
	"io"

	"github.com/n-creativesystem/short-url/pkg/domain/short"
)

type Service interface {
	GetURL(ctx context.Context, key string) (string, error)
	GenerateShortURL(ctx context.Context, url string, key, author string) (*short.ShortWithTimeStamp, error)
	GenerateQRCode(ctx context.Context, key string) (io.Reader, error)
	Remove(ctx context.Context, key, author string) error
	Update(ctx context.Context, key, author, url string) (*short.ShortWithTimeStamp, error)
	FindAll(ctx context.Context, author string) ([]short.ShortWithTimeStamp, error)
	FindByKeyAndAuthor(ctx context.Context, key, author string) (*short.ShortWithTimeStamp, error)
}
