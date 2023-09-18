package tracking

import (
	"context"
	"log/slog"
	"time"

	"github.com/n-creativesystem/short-url/pkg/utils/apps"
	"github.com/n-creativesystem/short-url/pkg/utils/logging"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
)

func NewTracerProvider(ctx context.Context, otelAgentAddr string) (*sdktrace.TracerProvider, func(), error) {
	traceClient := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(otelAgentAddr),
		otlptracegrpc.WithDialOption(grpc.WithBlock()),
	)
	exporter, err := otlptrace.New(ctx, traceClient)
	if err != nil {
		return nil, nil, err
	}
	serviceName := apps.ServiceName()
	version := apps.Version()
	environment := apps.TrackingEnvironment()
	r := NewResource(serviceName, version, environment)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(r),
	)
	pp := NewPropagator()
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(pp)

	cleanup := func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		if err := tp.ForceFlush(ctx); err != nil {
			slog.With(logging.WithErr(err)).Error(err.Error())
		}
		ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
		if err := tp.Shutdown(ctx2); err != nil {
			slog.With(logging.WithErr(err)).Error(err.Error())
		}
		cancel()
		cancel2()
	}
	return tp, cleanup, nil
}

func NewResource(serviceName string, version string, environment string) *resource.Resource {
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(serviceName),
		semconv.ServiceVersionKey.String(version),
		semconv.DeploymentEnvironmentKey.String(environment),
		attribute.String("environment", environment),
		attribute.String("env", environment),
	)
}

func NewPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}
