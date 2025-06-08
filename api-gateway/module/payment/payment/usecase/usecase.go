package usecase

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/telemetry"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/payment/v1"
	paymentSvc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/payment/payment/service"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
)

type (
	IPaymentUseCase interface {
		FindPaymentById(ctx context.Context, requestId string, request *paymentRpc.FindPaymentByIdRequest) (*paymentRpc.Payment, error)
		FindPaymentByUserIdAndStatus(ctx context.Context, requestId string, request *paymentRpc.FindPaymentByUserIdAndStatusRequest) (*paymentRpc.Payment, error)
	}

	paymentUseCase struct {
		paymentSvc              paymentSvc.IPaymentService
		rabbitMQ                rabbitmqInfrastructure.IRabbitMQInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  pkg.IZapLogger
	}
)

func NewPaymentUseCase(
	paymentSvc paymentSvc.IPaymentService,
	rabbitMQ rabbitmqInfrastructure.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkg.IZapLogger,
) IPaymentUseCase {
	return &paymentUseCase{
		paymentSvc:              paymentSvc,
		rabbitMQ:                rabbitMQ,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
