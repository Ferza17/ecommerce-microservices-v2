//go:build wireinject
// +build wireinject

package rabbitmq

import (
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
)

// ProvideRabbitMQInfrastructure wires dependencies for IRabbitMQInfrastructure.
func ProvideRabbitMQInfrastructure(logger logger.IZapLogger) IRabbitMQInfrastructure {
	wire.Build(NewRabbitMQInfrastructure)
	return NewRabbitMQInfrastructure(logger)
}
