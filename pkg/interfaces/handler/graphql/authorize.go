package graphql

import (
	"context"

	"github.com/n-creativesystem/short-url/pkg/interfaces/middleware/session"
	"github.com/n-creativesystem/short-url/pkg/utils/hash"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func authorize(ctx context.Context) (string, error) {
	var user string

	if v, ok := session.GetAuthUserWithContext(ctx); ok {
		user = v.Email
	}
	if user == "" {
		return "", gqlerror.Errorf("Invalid authorization")
	}
	return hash.Sum([]byte(user)), nil
}
