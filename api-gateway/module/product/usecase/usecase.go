package usecase

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/rabbitmq"
	rpcClientInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/service"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/pb"
	gatewayEventStoreUseCase "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/gatewayEventStore/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
)

type (
	IProductUseCase interface {
		CreateProduct(ctx context.Context, requestId string, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error)
		FindProductById(ctx context.Context, requestId string, req *pb.FindProductByIdRequest) (*pb.Product, error)
	}
	ProductUseCase struct {
		rpcClient                rpcClientInfrastructure.Service
		rabbitMQ                 rabbitmqInfrastructure.IRabbitMQInfrastructure
		gatewayEventStoreUseCase gatewayEventStoreUseCase.IGatewayEventStoreUseCase
		logger                   pkg.IZapLogger
	}
)

func NewProductUseCase(
	rpcClient rpcClientInfrastructure.Service,
	rabbitMQ rabbitmqInfrastructure.IRabbitMQInfrastructure,
	gatewayEventStoreUseCase gatewayEventStoreUseCase.IGatewayEventStoreUseCase,
	logger pkg.IZapLogger,
) IProductUseCase {
	return &ProductUseCase{
		rpcClient:                rpcClient,
		rabbitMQ:                 rabbitMQ,
		gatewayEventStoreUseCase: gatewayEventStoreUseCase,
		logger:                   logger,
	}
}
