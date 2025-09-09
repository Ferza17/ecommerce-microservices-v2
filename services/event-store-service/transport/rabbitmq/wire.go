//go:build wireinject
// +build wireinject

package rabbitmq

import (
	mongodbInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/mongodb"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/telemetry"
	eventConsumer "github.com/ferza17/ecommerce-microservices-v2/event-store-service/module/event/consumer"
	eventRepository "github.com/ferza17/ecommerce-microservices-v2/event-store-service/module/event/repository/mongodb"
	eventUseCase "github.com/ferza17/ecommerce-microservices-v2/event-store-service/module/event/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/logger"
	"github.com/google/wire"
)

func Provide() *Transport {
	wire.Build(
		logger.Set,
		// Infrastructure Layer
		mongodbInfrastructure.Set,
		rabbitmqInfrastructure.Set,
		telemetryInfrastructure.Set,
		// Repository Layer
		eventRepository.Set,
		// UseCase Layer
		eventUseCase.Set,
		// Consumer
		eventConsumer.Set,
		Set,
	)
	return nil
}
