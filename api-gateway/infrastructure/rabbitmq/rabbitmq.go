package rabbitmq

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/config"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
	"github.com/rabbitmq/amqp091-go"
)

type (
	IRabbitMQInfrastructure interface {
		Close() error
	}
	RabbitMQInfrastructure struct {
		amqpConn *amqp091.Connection
		amqpChan *amqp091.Channel
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

	amqpChannel, err := amqpConn.Channel()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create a channel: %v", err))
	}

	return &RabbitMQInfrastructure{
		amqpConn: amqpConn,
		amqpChan: amqpChannel,
		logger:   logger,
	}
}

func (c *RabbitMQInfrastructure) Close() error {
	if err := c.amqpConn.Close(); err != nil {
		c.logger.Error(fmt.Sprintf("Failed to close a connection: %v", err))
		return err
	}

	if err := c.amqpChan.Close(); err != nil {
		c.logger.Error(fmt.Sprintf("Failed to close a channel: %v", err))
		return err
	}

	return nil
}
