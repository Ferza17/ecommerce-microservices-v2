package telemetry

import (
	"context"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/context"
	"github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

func (t *telemetryInfrastructure) StartSpanFromRabbitMQHeader(ctx context.Context, headers amqp091.Table, fnName string) (context.Context, trace.Span) {
	var span trace.Span

	if headers != nil {
		if _, ok := headers[pkgContext.ContextKeyTracerparent].(string); ok {
			carrier := make(propagation.MapCarrier)
			for key, value := range headers {
				if strValue, ok := value.(string); ok {
					carrier.Set(key, strValue)
				}
			}

			ctx, span = t.StartSpanFromContext(t.extractSpanFromTextMapPropagator(ctx, carrier), fnName)
		} else {
			ctx, span = t.StartSpanFromContext(ctx, fnName)
		}
	} else {
		ctx, span = t.StartSpanFromContext(ctx, fnName)
	}

	span.SetAttributes(
		attribute.String("messaging.system", "rabbitmq"),
		attribute.String("messaging.operation", "consume"),
	)

	return ctx, span
}
