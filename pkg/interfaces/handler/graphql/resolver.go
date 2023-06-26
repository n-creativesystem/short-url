package graphql

import (
	oauth2client "github.com/n-creativesystem/short-url/pkg/service/oauth2_client"
	"github.com/n-creativesystem/short-url/pkg/service/short"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	oauth2clientSvc oauth2client.Service
	shortUrlSvc     short.Service
}

func NewResolver(
	oauth2client oauth2client.Service,
	shortUrlSvc short.Service,
) *Resolver {
	return &Resolver{
		oauth2clientSvc: oauth2client,
		shortUrlSvc:     shortUrlSvc,
	}
}
