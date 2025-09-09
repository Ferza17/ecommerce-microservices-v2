package telemetry

import (
	"context"
	"go.opentelemetry.io/otel/trace"
)

func (t *telemetryInfrastructure) StartSpanFromContext(ctx context.Context, fnName string) (context.Context, trace.Span) {
	return t.tracerProvider.Tracer(t.serviceName).Start(ctx, fnName)
}
