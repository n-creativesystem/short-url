//go:generate go run -mod=mod github.com/golang/mock/mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../mock/service/$GOPACKAGE/$GOFILE
package oauth2client

import (
	"context"

	"github.com/go-oauth2/oauth2/v4"
	oauth2client "github.com/n-creativesystem/short-url/pkg/domain/oauth2_client"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/tracking"
)

type RegisterResult struct {
	ClientId     string
	ClientSecret string
}

type Service interface {
	oauth2.ClientStore
	FindAll(ctx context.Context, user string) ([]oauth2client.Client, error)
	FindByID(ctx context.Context, id, user string) (*oauth2client.Client, error)
	RegisterClient(ctx context.Context, user, appName string) (RegisterResult, error)
	UpdateClient(ctx context.Context, id, user, appName string) error
	DeleteClient(ctx context.Context, user, clientId string) error
}

var (
	tracer = tracking.Tracer("oauth2_client_service")
)
