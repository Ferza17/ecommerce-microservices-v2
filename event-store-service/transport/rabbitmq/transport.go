package rabbitmq

import (
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/bootstrap"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/module/event/consumer"
	pkgLogger "github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/logger"
)

type (
	RabbitMQTransport struct {
		eventConsumer consumer.IEventConsumer
		logger        pkgLogger.IZapLogger
	}
)

func NewServer(dependency *bootstrap.Bootstrap) *RabbitMQTransport {
	return &RabbitMQTransport{
		logger:        dependency.Logger,
		eventConsumer: dependency.EventConsumer,
	}
}
