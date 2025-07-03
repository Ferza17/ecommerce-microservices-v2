package telemetry

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
	"github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

type (
	ITelemetryInfrastructure interface {
		Shutdown(ctx context.Context) error
		StartSpanFromContext(ctx context.Context, fnName string) (context.Context, trace.Span)
		StartSpanFromHttpRequest(r *http.Request, fnName string) (context.Context, trace.Span)
		StartSpanFromRabbitMQHeader(ctx context.Context, headers amqp091.Table, fnName string) (context.Context, trace.Span)
		StartSpanFromRpcMetadata(ctx context.Context, fnName string) (context.Context, trace.Span)

		InjectSpanToTextMapPropagator(ctx context.Context) propagation.MapCarrier
	}
	telemetryInfrastructure struct {
		logger         logger.IZapLogger
		tracerProvider *sdktrace.TracerProvider
		serviceName    string
	}
)

var Set = wire.NewSet(NewTelemetry)

func NewTelemetry(logger logger.IZapLogger) ITelemetryInfrastructure {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(
		jaeger.WithEndpoint(fmt.Sprintf("http://%s:%s/api/traces",
			config.Get().JaegerTelemetryHost,
			config.Get().JaegerTelemetryPort,
		)),
	))
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create a jaeger exporter: %v", err))
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(config.Get().UserServiceServiceName),
		)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	return &telemetryInfrastructure{
		logger:         logger,
		tracerProvider: tp,
		serviceName:    config.Get().UserServiceServiceName,
	}
}
