package graphql

import "github.com/n-creativesystem/short-url/pkg/infrastructure/tracking"

var (
	tracer = tracking.Tracer("graphql")
)
