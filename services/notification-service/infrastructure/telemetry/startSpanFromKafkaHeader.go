package telemetry

import (
	"context"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

func (t *telemetryInfrastructure) StartSpanFromKafkaHeader(ctx context.Context, headers []kafka.Header, fnName string) (context.Context, trace.Span) {
	var span trace.Span
	if headers != nil && len(headers) > 0 {
		for _, header := range headers {
			if strings.ToLower(header.Key) == strings.ToLower(pkgContext.ContextKeyTracerparent) {
				carrier := make(propagation.MapCarrier)
				carrier.Set(header.Key, string(header.Value))
				break
			}
			ctx, span = t.StartSpanFromContext(ctx, fnName)
		}
	} else {
		ctx, span = t.StartSpanFromContext(ctx, fnName)
	}

	span.SetAttributes(
		attribute.String("messaging.system", "kafka"),
		attribute.String("messaging.operation", "consume"),
	)

	return ctx, span
}
