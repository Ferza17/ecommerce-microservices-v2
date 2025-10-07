//go:build wireinject
// +build wireinject

package mongodb

import (
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
	"github.com/google/wire"
)

func ProvideMongoDBInfrastructure() IMongoDBInfrastructure {
	wire.Build(
		logger.Set,
		Set,
	)
	return nil
}
