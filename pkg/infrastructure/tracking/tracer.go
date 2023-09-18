package tracking

import (
	"fmt"

	"github.com/n-creativesystem/short-url/pkg/infrastructure/tracking/trace"
	"github.com/n-creativesystem/short-url/pkg/utils/apps"
	"go.opentelemetry.io/otel"
	oteltrace "go.opentelemetry.io/otel/trace"
)

func Tracer(name string, opts ...oteltrace.TracerOption) oteltrace.Tracer {
	t := otel.Tracer(fmt.Sprintf("%s/%s", apps.ServerRoot(), name), opts...)
	return trace.Tracer(t)
}
