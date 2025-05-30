package usecase

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/telemetry"
	eventRpc "github.com/ferza17/ecommerce-microservices-v2/event-store-service/model/rpc/gen/event/v1"
	eventRepository "github.com/ferza17/ecommerce-microservices-v2/event-store-service/module/event/repository/mongodb"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg"
)

type (
	IEventUseCase interface {
		CreateEventStore(ctx context.Context, requestId string, req *eventRpc.EventStore) (*eventRpc.CreateEventStoreResponse, error)
	}

	eventUseCase struct {
		eventRepository         eventRepository.IEventRepository
		rabbitMQInfrastructure  rabbitmqInfrastructure.IRabbitMQInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  pkg.IZapLogger
	}
)

func NewEventStoreUseCase(
	eventStoreRepository eventRepository.IEventRepository,
	rabbitMQInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkg.IZapLogger,
) IEventUseCase {
	return &eventUseCase{
		eventRepository:         eventStoreRepository,
		rabbitMQInfrastructure:  rabbitMQInfrastructure,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
