package usecase

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/telemetry"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/payment/v1"
	paymentSvc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/payment/payment/service"
	providerSvc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/payment/provider/service"
	productSvc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/product/service"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
)

type (
	IPaymentUseCase interface {
		FindPaymentById(ctx context.Context, requestId string, request *paymentRpc.FindPaymentByIdRequest) (*paymentRpc.Payment, error)
		FindPaymentByUserIdAndStatus(ctx context.Context, requestId string, request *paymentRpc.FindPaymentByUserIdAndStatusRequest) (*paymentRpc.Payment, error)

		CretePayment(ctx context.Context, requestId string, request *paymentRpc.CreatePaymentRequest) error
	}

	paymentUseCase struct {
		paymentSvc              paymentSvc.IPaymentService
		productSvc              productSvc.IProductService
		providerSvc             providerSvc.IPaymentProviderService
		rabbitMQ                rabbitmqInfrastructure.IRabbitMQInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  pkg.IZapLogger
	}
)

func NewPaymentUseCase(
	paymentSvc paymentSvc.IPaymentService,
	productSvc productSvc.IProductService,
	providerSvc providerSvc.IPaymentProviderService,
	rabbitMQ rabbitmqInfrastructure.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkg.IZapLogger,
) IPaymentUseCase {
	return &paymentUseCase{
		paymentSvc:              paymentSvc,
		rabbitMQ:                rabbitMQ,
		telemetryInfrastructure: telemetryInfrastructure,
		productSvc:              productSvc,
		providerSvc:             providerSvc,
		logger:                  logger,
	}
}
