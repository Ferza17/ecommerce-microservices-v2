package consumer

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IPaymentConsumer interface {
		SnapshotPaymentsPaymentOrderCreated(ctx context.Context, message *kafka.Message) error
		CompensateSnapshotPaymentsPaymentOrderCreated(ctx context.Context, message *kafka.Message) error
		ConfirmSnapshotPaymentsPaymentOrderCreated(ctx context.Context, message *kafka.Message) error

		SnapshotPaymentsPaymentOrderCancelledDelayed(ctx context.Context, message *kafka.Message) error
	}

	paymentConsumer struct {
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
	telemetryInfrastructure telemetry.ITelemetryInfrastructure,
	paymentUseCase usecase.IPaymentUseCase,
	logger logger.IZapLogger,
) IPaymentConsumer {
	return &paymentConsumer{
		telemetryInfrastructure: telemetryInfrastructure,
		paymentUseCase:          paymentUseCase,
		logger:                  logger,
	}
}
