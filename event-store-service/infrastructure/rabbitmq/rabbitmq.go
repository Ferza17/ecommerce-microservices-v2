package rabbitmq

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg"
	"github.com/rabbitmq/amqp091-go"
)

type (
	IRabbitMQInfrastructure interface {
		Close() error
	}
	RabbitMQInfrastructure struct {
		amqpConn *amqp091.Connection
		logger   pkg.IZapLogger
	}
)

func NewRabbitMQInfrastructure(logger pkg.IZapLogger) IRabbitMQInfrastructure {
	amqpConn, err := amqp091.Dial(
		fmt.Sprintf("amqp://%s:%s@%s:%s/",
			config.Get().RabbitMQUsername,
			config.Get().RabbitMQPassword,
			config.Get().RabbitMQHost,
			config.Get().RabbitMQPort,
		))
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to connect to RabbitMQ: %v", err))
	}

	return &RabbitMQInfrastructure{
		amqpConn: amqpConn,
		logger:   logger,
	}
}

func (c *RabbitMQInfrastructure) Close() error {
	if err := c.amqpConn.Close(); err != nil {
		c.logger.Error(fmt.Sprintf("Failed to close a connection: %v", err))
		return err
	}

	return nil
}
