//go:build wireinject
// +build wireinject

package kafka

import (
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
	"github.com/google/wire"
)

func ProvideKafkaInfrastructure() IKafkaInfrastructure {
	wire.Build(
		logger.Set,
		telemetryInfrastructure.Set,
		Set,
	)

	return nil
}
