//go:generate go run -mod=mod github.com/golang/mock/mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../mock/repository/$GOPACKAGE/$GOFILE
package oauth2client

import (
	"context"

	"github.com/go-oauth2/oauth2/v4"
)

type Repository interface {
	GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error)
	Delete(ctx context.Context, user, id string) error
	Create(ctx context.Context, client *Client) error
	Find(ctx context.Context, user string) ([]Client, error)
	FindByID(ctx context.Context, id, user string) (*Client, error)
	Update(ctx context.Context, id, user, appName string) error
}
