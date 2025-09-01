package consumer

import (
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/telemetry"
	eventUseCase "github.com/ferza17/ecommerce-microservices-v2/event-store-service/module/event/usecase"
	pkgLogger "github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IEventConsumer interface {
	}

	eventConsumer struct {
		eventUseCase            eventUseCase.IEventUseCase
		rabbitMQInfrastructure  rabbitmqInfrastructure.IRabbitMQInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  pkgLogger.IZapLogger
	}
)

var Set = wire.NewSet(NewEventConsumer)

func NewEventConsumer(
	rabbitMQInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	eventUseCase eventUseCase.IEventUseCase,
	logger pkgLogger.IZapLogger,
) IEventConsumer {
	return &eventConsumer{
		eventUseCase:            eventUseCase,
		logger:                  logger,
		rabbitMQInfrastructure:  rabbitMQInfrastructure,
		telemetryInfrastructure: telemetryInfrastructure,
	}
}
