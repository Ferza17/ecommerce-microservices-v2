package telemetry

import (
	"context"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/context"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/metadata"
)

func (t *telemetryInfrastructure) StartSpanFromRpcMetadata(ctx context.Context, fnName string) (context.Context, trace.Span) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}

	var span trace.Span

	if traceparents := md.Get(pkgContext.ContextKeyTracerparent); len(traceparents) > 0 {
		carrier := make(propagation.MapCarrier)
		carrier.Set(pkgContext.ContextKeyTracerparent, traceparents[0])

		ctx, span = t.StartSpanFromContext(
			t.extractSpanFromTextMapPropagator(ctx, carrier),
			fnName,
		)
	} else {
		ctx, span = t.StartSpanFromContext(ctx, fnName)
	}

	return ctx, span
}
