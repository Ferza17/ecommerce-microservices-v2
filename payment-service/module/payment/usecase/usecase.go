package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/payment/v1"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/repository"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
)

type (
	IPaymentUseCase interface {
		CreatePayment(ctx context.Context, requestId string, request *paymentRpc.CreatePaymentRequest) error

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
