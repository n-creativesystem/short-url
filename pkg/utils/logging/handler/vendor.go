package handler

import (
	"context"
	"log/slog"
	"strconv"
	"strings"

	"github.com/n-creativesystem/short-url/pkg/utils"
	"github.com/n-creativesystem/short-url/pkg/utils/apps"
	"go.opentelemetry.io/otel/trace"
)

func spanVendorConnector(ctx context.Context) []slog.Attr {
	vendor := utils.Getenv("APM_VENDOR", "")
	if strings.EqualFold(vendor, "datadog") {
		return datadog(ctx)
	} else {
		return default_(ctx)
	}
}

func default_(ctx context.Context) []slog.Attr {
	span := trace.SpanFromContext(ctx)
	if span == nil {
		return nil
	}
	return []slog.Attr{
		slog.String("trace_id", span.SpanContext().TraceID().String()),
		slog.String("span_id", span.SpanContext().SpanID().String()),
		slog.String("service", apps.ServiceName()),
		slog.String("env", apps.TrackingEnvironment()),
		slog.String("version", apps.Version()),
	}
}

func datadog(ctx context.Context) []slog.Attr {
	span := trace.SpanFromContext(ctx)
	if span == nil {
		return nil
	}
	return []slog.Attr{
		slog.String("dd.trace_id", otelTraceIDToDDTrace(span.SpanContext().TraceID().String())),
		slog.String("dd.span_id", otelTraceIDToDDTrace(span.SpanContext().SpanID().String())),
		slog.String("dd.service", apps.ServiceName()),
		slog.String("dd.env", apps.TrackingEnvironment()),
		slog.String("dd.version", apps.Version()),
	}
}

func otelTraceIDToDDTrace(id string) string {
	if len(id) < 16 {
		return ""
	}
	if len(id) > 16 {
		id = id[16:]
	}
	intValue, err := strconv.ParseUint(id, 16, 64)
	if err != nil {
		return ""
	}
	return strconv.FormatUint(intValue, 10)
}
