package consumer

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/telemetry"
	eventUseCase "github.com/ferza17/ecommerce-microservices-v2/event-store-service/module/event/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg"
)

type (
	IEventConsumer interface {
		EventCreated(ctx context.Context) error
	}

	eventConsumer struct {
		eventUseCase            eventUseCase.IEventUseCase
		rabbitMQInfrastructure  rabbitmqInfrastructure.IRabbitMQInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  pkg.IZapLogger
	}
)

func NewEventConsumer(
	rabbitMQInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	eventUseCase eventUseCase.IEventUseCase,
	logger pkg.IZapLogger,
) IEventConsumer {
	return &eventConsumer{
		eventUseCase:            eventUseCase,
		logger:                  logger,
		rabbitMQInfrastructure:  rabbitMQInfrastructure,
		telemetryInfrastructure: telemetryInfrastructure,
	}
}
