package consumer

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
)

type (
	IPaymentConsumer interface {
		PaymentOrderCreated(ctx context.Context) error
		PaymentOrderDelayedCancelled(ctx context.Context) error
	}

	paymentConsumer struct {
		rabbitmq                rabbitmq.IRabbitMQInfrastructure
		telemetryInfrastructure telemetry.ITelemetryInfrastructure
		paymentUseCase          usecase.IPaymentUseCase
		logger                  logger.IZapLogger
	}
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
