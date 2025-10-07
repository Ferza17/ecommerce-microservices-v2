package usecase

import (
	"context"

	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/kafka"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/postgres"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	pbEvent "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/event"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	eventMongoDBRepository "github.com/ferza17/ecommerce-microservices-v2/product-service/module/event/repository/mongodb"
	eventUseCase "github.com/ferza17/ecommerce-microservices-v2/product-service/module/event/usecase"
	productElasticsearchRepository "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/repository/elasticsearch"
	productRepo "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/repository/postgres"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"

	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
)

type (
	IProductUseCase interface {
		FindProductById(ctx context.Context, requestId string, req *productRpc.FindProductByIdRequest) (*productRpc.Product, error)
		FindProductsWithPagination(ctx context.Context, requestId string, req *productRpc.FindProductsWithPaginationRequest) (*productRpc.FindProductsWithPaginationResponse, error)

		CreateProduct(ctx context.Context, requestId string, req *productRpc.CreateProductRequest) (*empty.Empty, error)
		ConfirmCreateProduct(ctx context.Context, requestId string, req *pbEvent.ReserveEvent) error
		CompensateCreateProduct(ctx context.Context, requestId string, req *pbEvent.ReserveEvent) error

		UpdateProductById(ctx context.Context, requestId string, req *productRpc.UpdateProductByIdRequest) (*empty.Empty, error)
		ConfirmUpdateProductById(ctx context.Context, requestId string, req *pbEvent.ReserveEvent) error
		CompensateUpdateProductById(ctx context.Context, requestId string, req *pbEvent.ReserveEvent) error

		DeleteProductById(ctx context.Context, requestId string, req *productRpc.DeleteProductByIdRequest) (*empty.Empty, error)
		ConfirmDeleteProductById(ctx context.Context, requestId string, req *pbEvent.ReserveEvent) error
		CompensateDeleteProductById(ctx context.Context, requestId string, req *pbEvent.ReserveEvent) error
	}

	productUseCase struct {
		postgres                       *postgres.PostgresSQL
		productPgsqlRepository         productRepo.IProductPostgresqlRepository
		eventMongoDBRepository         eventMongoDBRepository.IEventMongoRepository
		kafkaInfrastructure            kafka.IKafkaInfrastructure
		productElasticsearchRepository productElasticsearchRepository.IProductElasticsearchRepository
		telemetryInfrastructure        telemetryInfrastructure.ITelemetryInfrastructure
		eventUseCase                   eventUseCase.IEventUseCase
		logger                         logger.IZapLogger
	}
)

var Set = wire.NewSet(NewProductUseCase)

func NewProductUseCase(
	postgres *postgres.PostgresSQL,
	productPgsqlRepository productRepo.IProductPostgresqlRepository,
	eventMongoDBRepository eventMongoDBRepository.IEventMongoRepository,
	kafkaInfrastructure kafka.IKafkaInfrastructure,
	productElasticsearchRepository productElasticsearchRepository.IProductElasticsearchRepository,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	eventUseCase eventUseCase.IEventUseCase,
	logger logger.IZapLogger,
) IProductUseCase {
	return &productUseCase{
		postgres:                       postgres,
		productPgsqlRepository:         productPgsqlRepository,
		eventMongoDBRepository:         eventMongoDBRepository,
		kafkaInfrastructure:            kafkaInfrastructure,
		telemetryInfrastructure:        telemetryInfrastructure,
		productElasticsearchRepository: productElasticsearchRepository,
		eventUseCase:                   eventUseCase,
		logger:                         logger,
	}
}
