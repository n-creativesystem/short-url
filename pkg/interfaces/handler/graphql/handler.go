package graphql

import (
	"context"
	"log/slog"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func GraphQLHandler(resolver ResolverRoot) gin.HandlerFunc {
	h := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: resolver}))
	h.SetRecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
		slog.With("err", err).ErrorContext(ctx, "Graphql request recover.")
		return &gqlerror.Error{
			Message: "Internal server error.",
			Extensions: map[string]interface{}{
				"code": 500,
			},
		}
	})
	return gin.WrapH(h)
}

func GraphQLPlayGroundHandler(endpoint string) gin.HandlerFunc {
	h := playground.Handler("Graphql", endpoint)
	return gin.WrapH(h)
}
