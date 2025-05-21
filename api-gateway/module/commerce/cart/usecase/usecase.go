package usecase

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/pb"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
)

type (
	ICartUseCase interface {
		CreateCart(ctx context.Context, requestId string, req *pb.CreateCartItemRequest) (*pb.CreateCartItemResponse, error)
		UpdateCartItemById(ctx context.Context, requestId string, req *pb.UpdateCartItemByIdRequest) (*pb.UpdateCartItemByIdResponse, error)
	}

	CartUseCase struct {
		rabbitMQ                rabbitmqInfrastructure.IRabbitMQInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  pkg.IZapLogger
	}
)

func NewCartUseCase(
	rabbitMQ rabbitmqInfrastructure.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkg.IZapLogger,
) ICartUseCase {
	return &CartUseCase{
		rabbitMQ:                rabbitMQ,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
