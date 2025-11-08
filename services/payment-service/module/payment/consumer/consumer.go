package consumer

import (
	"context"

	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	pbEvent "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/event"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IPaymentConsumer interface {
		SnapshotPaymentsPaymentOrderCreated(ctx context.Context, message *pbEvent.EventEnvelope) error
		SnapshotPaymentsPaymentOrderCancelledDelayed(ctx context.Context, message *pbEvent.EventEnvelope) error
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
