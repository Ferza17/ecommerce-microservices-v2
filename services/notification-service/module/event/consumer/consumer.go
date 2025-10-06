package consumer

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	kafkaInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/kafka"
	eventUseCase "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/event/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IEventConsumer interface {
		DlqSinkMongoEventsNotificationEventStores(ctx context.Context, message *kafka.Message) error
	}

	eventConsumer struct {
		kafkaInfrastructure kafkaInfrastructure.IKafkaInfrastructure
		logger              logger.IZapLogger
		eventUseCase        eventUseCase.IEventUseCase
	}
)

var Set = wire.NewSet(NewEventConsumer)

func NewEventConsumer(
	kafkaInfrastructure kafkaInfrastructure.IKafkaInfrastructure,
	logger logger.IZapLogger,
	eventUseCase eventUseCase.IEventUseCase,
) IEventConsumer {
	return &eventConsumer{
		kafkaInfrastructure: kafkaInfrastructure,
		logger:              logger,
		eventUseCase:        eventUseCase,
	}
}
