package telemetry

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
)

type (
	ITelemetryInfrastructure interface {
		Close(ctx context.Context) error
		Tracer(ctx context.Context, fnName string) (context.Context, trace.Span)
	}
	telemetryInfrastructure struct {
		logger         pkg.IZapLogger
		tracerProvider *sdktrace.TracerProvider
	}
)

func NewTelemetry(logger pkg.IZapLogger) ITelemetryInfrastructure {
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
			semconv.ServiceNameKey.String(config.Get().ServiceName),
		)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return &telemetryInfrastructure{
		logger:         logger,
		tracerProvider: tp,
	}
}

func (t *telemetryInfrastructure) Close(ctx context.Context) error {
	if err := t.tracerProvider.Shutdown(ctx); err != nil {
		t.logger.Error(fmt.Sprintf("Failed to shutdown tracer provider: %v", err))
		return err
	}
	return nil
}

func (t *telemetryInfrastructure) Tracer(ctx context.Context, fnName string) (context.Context, trace.Span) {
	return t.tracerProvider.Tracer(config.Get().ServiceName).Start(ctx, fnName)
}
