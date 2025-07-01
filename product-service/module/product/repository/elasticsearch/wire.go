//go:build wireinject
// +build wireinject

package elasticsearch

import (
	elasticsearchInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/elasticsearch"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
	"github.com/google/wire"
)

func ProvideProductElasticsearchRepository() IProductElasticsearchRepository {
	wire.Build(
		elasticsearchInfrastructure.Set,
		telemetryInfrastructure.Set,
		logger.Set,
		Set,
	)

	return nil
}
