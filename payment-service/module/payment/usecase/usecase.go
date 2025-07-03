package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/repository"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IPaymentUseCase interface {
		CreatePayment(ctx context.Context, requestId string, request *paymentRpc.CreatePaymentRequest) error
		PaymentOrderDelayedCancelled(ctx context.Context, requestId string, request *paymentRpc.PaymentOrderDelayedCancelledRequest) error

		FindPaymentById(ctx context.Context, requestId string, request *paymentRpc.FindPaymentByIdRequest) (*paymentRpc.Payment, error)
		FindPaymentByUserIdAndStatus(ctx context.Context, requestId string, request *paymentRpc.FindPaymentByUserIdAndStatusRequest) (*paymentRpc.Payment, error)
	}

	paymentUseCase struct {
		paymentRepository       repository.IPaymentRepository
		rabbitmqInfrastructure  rabbitmq.IRabbitMQInfrastructure
		telemetryInfrastructure telemetry.ITelemetryInfrastructure
		logger                  logger.IZapLogger
	}
)

// Set is a Wire provider set for Payment use case dependencies
var Set = wire.NewSet(
	NewPaymentUseCase,
)

func NewPaymentUseCase(
	paymentRepository repository.IPaymentRepository,
	rabbitmqInfrastructure rabbitmq.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetry.ITelemetryInfrastructure,
	logger logger.IZapLogger,
) IPaymentUseCase {
	return &paymentUseCase{
		paymentRepository:       paymentRepository,
		rabbitmqInfrastructure:  rabbitmqInfrastructure,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
