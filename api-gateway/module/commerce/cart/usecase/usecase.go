package usecase

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/rabbitmq"
	rpcClientInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/service"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/pb"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
)

type (
	ICartUseCase interface {
		CreateCart(ctx context.Context, requestId string, req *pb.CreateCartItemRequest) (*pb.CreateCartItemResponse, error)
		UpdateCartItemById(ctx context.Context, requestId string, req *pb.UpdateCartItemByIdRequest) (*pb.UpdateCartItemByIdResponse, error)
	}

	CartUseCase struct {
		rpcClient rpcClientInfrastructure.IService
		rabbitMQ  rabbitmqInfrastructure.IRabbitMQInfrastructure
		logger    pkg.IZapLogger
	}
)

func NewCartUseCase(
	rpcClient rpcClientInfrastructure.IService,
	rabbitMQ rabbitmqInfrastructure.IRabbitMQInfrastructure,
	logger pkg.IZapLogger,
) ICartUseCase {
	return &CartUseCase{
		rpcClient: rpcClient,
		rabbitMQ:  rabbitMQ,
		logger:    logger,
	}
}
