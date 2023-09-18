package datadog

import (
	"github.com/n-creativesystem/short-url/pkg/utils"
	"go.opentelemetry.io/otel"
	ddotel "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/opentelemetry"
)

func IsEnable() bool {
	return utils.GetBoolEnv("DD_TRACE_ENABLED")
}

func APM() func() {
	tp := ddotel.NewTracerProvider()
	otel.SetTracerProvider(tp)
	return func() { _ = tp.Shutdown() }
}
