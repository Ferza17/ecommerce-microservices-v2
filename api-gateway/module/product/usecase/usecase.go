package usecase

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/pb"
	productService "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/product/service"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
)

type (
	IProductUseCase interface {
		CreateProduct(ctx context.Context, requestId string, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error)
		FindProductById(ctx context.Context, requestId string, req *pb.FindProductByIdRequest) (*pb.Product, error)
		FindProductsWithPagination(ctx context.Context, requestId string, req *pb.FindProductsWithPaginationRequest) (*pb.FindProductsWithPaginationResponse, error)
	}
	ProductUseCase struct {
		productService          productService.IProductService
		rabbitMQ                rabbitmqInfrastructure.IRabbitMQInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  pkg.IZapLogger
	}
)

func NewProductUseCase(
	productService productService.IProductService,
	rabbitMQ rabbitmqInfrastructure.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkg.IZapLogger,
) IProductUseCase {
	return &ProductUseCase{
		productService:          productService,
		rabbitMQ:                rabbitMQ,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
