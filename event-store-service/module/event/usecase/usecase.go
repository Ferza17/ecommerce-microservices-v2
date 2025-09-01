package usecase

import (
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/telemetry"
	eventRepository "github.com/ferza17/ecommerce-microservices-v2/event-store-service/module/event/repository/mongodb"
	pkgLogger "github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IEventUseCase interface {
	}

	eventUseCase struct {
		eventRepository         eventRepository.IEventRepository
		rabbitMQInfrastructure  rabbitmqInfrastructure.IRabbitMQInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  pkgLogger.IZapLogger
	}
)

var Set = wire.NewSet(NewEventStoreUseCase)

func NewEventStoreUseCase(
	eventStoreRepository eventRepository.IEventRepository,
	rabbitMQInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkgLogger.IZapLogger,
) IEventUseCase {
	return &eventUseCase{
		eventRepository:         eventStoreRepository,
		rabbitMQInfrastructure:  rabbitMQInfrastructure,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
