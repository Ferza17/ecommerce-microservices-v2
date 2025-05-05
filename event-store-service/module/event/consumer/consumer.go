package consumer

import (
	"context"
	eventUseCase "github.com/ferza17/ecommerce-microservices-v2/event-store-service/module/event/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg"
	"github.com/rabbitmq/amqp091-go"
)

type (
	IEventConsumer interface {
		EventCreated(ctx context.Context) error
		Close(ctx context.Context) error
	}

	eventConsumer struct {
		amqpChannel  *amqp091.Channel
		eventUseCase eventUseCase.IEventUseCase
		logger       pkg.IZapLogger
	}
)

func NewEventConsumer(amqpChannel *amqp091.Channel, eventUseCase eventUseCase.IEventUseCase, logger pkg.IZapLogger) IEventConsumer {
	return &eventConsumer{
		amqpChannel:  amqpChannel,
		eventUseCase: eventUseCase,
		logger:       logger,
	}
}

func (c *eventConsumer) Close(ctx context.Context) error {
	return c.amqpChannel.Close()
}
