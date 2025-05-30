package usecase

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/product/v1"
	productElasticsearchRepository "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/repository/elasticsearch"
	productRepo "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/repository/postgresql"

	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg"
)

type (
	IProductUseCase interface {
		FindProductById(ctx context.Context, requestId string, req *productRpc.FindProductByIdRequest) (*productRpc.Product, error)
		FindProductsWithPagination(ctx context.Context, requestId string, req *productRpc.FindProductsWithPaginationRequest) (*productRpc.FindProductsWithPaginationResponse, error)
		CreateProduct(ctx context.Context, requestId string, req *productRpc.CreateProductRequest) (*productRpc.CreateProductResponse, error)
		UpdateProductById(ctx context.Context, requestId string, req *productRpc.UpdateProductByIdRequest) (*productRpc.Product, error)
		DeleteProductById(ctx context.Context, requestId string, req *productRpc.DeleteProductByIdRequest) (*productRpc.DeleteProductByIdResponse, error)
	}

	productUseCase struct {
		productPgsqlRepository         productRepo.IProductPostgresqlRepository
		productElasticsearchRepository productElasticsearchRepository.IProductElasticsearchRepository
		rabbitmqInfrastructure         rabbitmqInfrastructure.IRabbitMQInfrastructure
		telemetryInfrastructure        telemetryInfrastructure.ITelemetryInfrastructure
		logger                         pkg.IZapLogger
	}
)

func NewProductUseCase(
	productPgsqlRepository productRepo.IProductPostgresqlRepository,
	rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	productElasticsearchRepository productElasticsearchRepository.IProductElasticsearchRepository,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkg.IZapLogger,
) IProductUseCase {
	return &productUseCase{
		productPgsqlRepository:         productPgsqlRepository,
		rabbitmqInfrastructure:         rabbitmqInfrastructure,
		telemetryInfrastructure:        telemetryInfrastructure,
		productElasticsearchRepository: productElasticsearchRepository,
		logger:                         logger,
	}
}
