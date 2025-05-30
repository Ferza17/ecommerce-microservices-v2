package usecase

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/telemetry"
	commerceRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/commerce/v1"
	commerceCartService "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/commerce/cart/service"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
)

type (
	ICartUseCase interface {
		CreateCart(ctx context.Context, requestId string, req *commerceRpc.CreateCartItemRequest) (*commerceRpc.CreateCartItemResponse, error)
		UpdateCartItemById(ctx context.Context, requestId string, req *commerceRpc.UpdateCartItemByIdRequest) (*commerceRpc.UpdateCartItemByIdResponse, error)
	}

	CartUseCase struct {
		rabbitMQ                rabbitmqInfrastructure.IRabbitMQInfrastructure
		commerceCartService     commerceCartService.ICommerceCartService
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  pkg.IZapLogger
	}
)

func NewCartUseCase(
	rabbitMQ rabbitmqInfrastructure.IRabbitMQInfrastructure,
	commerceCartService commerceCartService.ICommerceCartService,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkg.IZapLogger,
) ICartUseCase {
	return &CartUseCase{
		rabbitMQ:                rabbitMQ,
		commerceCartService:     commerceCartService,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
