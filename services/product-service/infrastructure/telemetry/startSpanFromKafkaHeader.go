package telemetry

import (
	"context"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

func (t *telemetryInfrastructure) StartSpanFromKafkaHeader(ctx context.Context, headers []kafka.Header, fnName string) (context.Context, trace.Span) {
	var span trace.Span
	if headers != nil && len(headers) > 0 {
		carrier := make(propagation.MapCarrier)
		for _, header := range headers {
			if strings.ToLower(header.Key) == strings.ToLower(pkgContext.ContextKeyTracerparent) {
				carrier.Set(header.Key, string(header.Value))
				break
			}
		}
		ctx, span = t.StartSpanFromContext(t.extractSpanFromTextMapPropagator(ctx, carrier), fnName)
	} else {
		ctx, span = t.StartSpanFromContext(ctx, fnName)
	}

	span.SetAttributes(
		attribute.String("messaging.system", "kafka"),
		attribute.String("messaging.operation", "consume"),
	)

	return ctx, span
}
