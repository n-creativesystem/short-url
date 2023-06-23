//go:generate go run -mod=mod github.com/golang/mock/mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../mock/repository/$GOPACKAGE/$GOFILE
package oauth2token

import (
	"github.com/go-oauth2/oauth2/v4"
)

type Repository interface {
	Close()
	oauth2.TokenStore
}
