package elasticsearch

import (
	elasticsearchInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/elasticsearch"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg"
)

type (
	IProductElasticsearchRepository interface{}
	productElasticsearchRepository  struct {
		elasticsearchInfrastructure elasticsearchInfrastructure.IElasticsearchInfrastructure
		logger                      pkg.IZapLogger
	}
)

func NewProductElasticsearchRepository(
	elasticsearchInfrastructure elasticsearchInfrastructure.IElasticsearchInfrastructure,
	logger pkg.IZapLogger,
) IProductElasticsearchRepository {
	return &productElasticsearchRepository{
		elasticsearchInfrastructure: elasticsearchInfrastructure,
		logger:                      logger,
	}
}
