package elasticsearch

import (
	"context"
	elasticsearchInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/elasticsearch"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/orm"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/product/v1"

	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg"
)

type (
	IProductElasticsearchRepository interface {
		BulkCreateProduct(ctx context.Context, products []*orm.Product) error
		FindProductById(ctx context.Context, requestId string, id string) (*orm.Product, error)
		FindProductsWithPagination(ctx context.Context, requestId string, request *productRpc.FindProductsWithPaginationRequest) (products []*orm.Product, total int64, err error)
	}
	productElasticsearchRepository struct {
		elasticsearchInfrastructure elasticsearchInfrastructure.IElasticsearchInfrastructure
		telemetryInfrastructure     telemetryInfrastructure.ITelemetryInfrastructure
		logger                      pkg.IZapLogger
	}
)

const productIndex = "products"

func NewProductElasticsearchRepository(
	elasticsearchInfrastructure elasticsearchInfrastructure.IElasticsearchInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkg.IZapLogger,
) IProductElasticsearchRepository {
	return &productElasticsearchRepository{
		elasticsearchInfrastructure: elasticsearchInfrastructure,
		telemetryInfrastructure:     telemetryInfrastructure,
		logger:                      logger,
	}
}
