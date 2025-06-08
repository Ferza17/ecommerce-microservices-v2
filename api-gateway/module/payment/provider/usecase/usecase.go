package usecase

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/telemetry"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/payment/v1"
	paymentProviderSvc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/payment/provider/service"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
)

type (
	IPaymentProviderUseCase interface {
		FindPaymentProviders(ctx context.Context, requestId string, request *paymentRpc.FindPaymentProvidersRequest) (*paymentRpc.FindPaymentProvidersResponse, error)
	}

	paymentProviderUseCase struct {
		paymentProviderSvc      paymentProviderSvc.IPaymentProviderService
		rabbitMQ                rabbitmqInfrastructure.IRabbitMQInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  pkg.IZapLogger
	}
)

func NewProviderUseCase(
	paymentProviderSvc paymentProviderSvc.IPaymentProviderService,
	rabbitMQ rabbitmqInfrastructure.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkg.IZapLogger,
) IPaymentProviderUseCase {
	return &paymentProviderUseCase{
		paymentProviderSvc:      paymentProviderSvc,
		rabbitMQ:                rabbitMQ,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
