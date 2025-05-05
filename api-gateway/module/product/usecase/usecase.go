package usecase

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/rabbitmq"
	rpcClientInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/service"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/pb"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
)

type (
	IProductUseCase interface {
		CreateProduct(ctx context.Context, requestId string, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error)
		FindProductById(ctx context.Context, requestId string, req *pb.FindProductByIdRequest) (*pb.Product, error)
	}
	ProductUseCase struct {
		rpcClient rpcClientInfrastructure.IService
		rabbitMQ  rabbitmqInfrastructure.IRabbitMQInfrastructure
		logger    pkg.IZapLogger
	}
)

func NewProductUseCase(
	rpcClient rpcClientInfrastructure.IService,
	rabbitMQ rabbitmqInfrastructure.IRabbitMQInfrastructure,
	logger pkg.IZapLogger,
) IProductUseCase {
	return &ProductUseCase{
		rpcClient: rpcClient,
		rabbitMQ:  rabbitMQ,
		logger:    logger,
	}
}
