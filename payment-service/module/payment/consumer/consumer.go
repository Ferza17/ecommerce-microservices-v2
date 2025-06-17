package consumer

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IPaymentConsumer interface {
		PaymentOrderCreated(ctx context.Context) error
		PaymentOrderDelayedCancelled(ctx context.Context) error

		Close() error
	}

	paymentConsumer struct {
		rabbitmq                rabbitmq.IRabbitMQInfrastructure
		telemetryInfrastructure telemetry.ITelemetryInfrastructure
		paymentUseCase          usecase.IPaymentUseCase
		logger                  logger.IZapLogger
	}
)

// Set is a Wire provider set for Payment consumer dependencies
var Set = wire.NewSet(
	NewPaymentConsumer, // Wire will automatically infer the binding for IPaymentConsumer
)

// NewPaymentConsumer creates a new instance of IPaymentConsumer.
func NewPaymentConsumer(
	rabbitmq rabbitmq.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetry.ITelemetryInfrastructure,
	paymentUseCase usecase.IPaymentUseCase,
	logger logger.IZapLogger,
) IPaymentConsumer {
	return &paymentConsumer{
		rabbitmq:                rabbitmq,
		telemetryInfrastructure: telemetryInfrastructure,
		paymentUseCase:          paymentUseCase,
		logger:                  logger,
	}
}

func (c *paymentConsumer) Close() error {
	if err := c.rabbitmq.Close(); err != nil {
		c.logger.Error(fmt.Sprintf("Failed to close a connection: %v", err))
		return err
	}

	return nil
}
