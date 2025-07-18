package telemetry

import (
	"context"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

func (t *telemetryInfrastructure) StartSpanFromHttpRequest(r *http.Request, fnName string) (context.Context, trace.Span) {
	var span trace.Span
	ctx := r.Context()

	if traceparent := r.Header.Get(pkgContext.ContextKeyTracerparent); traceparent != "" {
		carrier := propagation.HeaderCarrier(r.Header)
		ctx, span = t.StartSpanFromContext(otel.GetTextMapPropagator().Extract(ctx, carrier), fnName)
	} else {
		ctx, span = t.StartSpanFromContext(ctx, fnName)
	}

	span.SetAttributes(
		attribute.String("http.method", r.Method),
		attribute.String("http.url", r.URL.String()),
		attribute.String("http.scheme", r.URL.Scheme),
	)

	return ctx, span
}
