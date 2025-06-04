//go:build wireinject
// +build wireinject

package rabbitmq

import (
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
)

// ProvideRabbitMQServer initializes a rabbitMQServer using Wire.
func ProvideRabbitMQServer() IRabbitMQServer {
	wire.Build(
		rabbitmq.ProvideRabbitMQInfrastructure, // Provides IRabbitMQInfrastructure
		logger.ProvideLogger,                   // Provides IZapLogger
		NewRabbitMQServer,                      // Combines dependencies into IRabbitMQServer
	)
	return nil
}
