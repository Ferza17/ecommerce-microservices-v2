package connector

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

type (
	IAmqpConnector interface {
		Close() error
	}
	AmqpConnector struct {
		amqpConn *amqp091.Connection
		amqpChan *amqp091.Channel
	}
)

func NewAmqpConnector() IAmqpConnector {
	amqpConn, err := amqp091.Dial(
		fmt.Sprintf("amqp://%s:%s@%s:%s/",
			config.Get().RabbitMQUsername,
			config.Get().RabbitMQPassword,
			config.Get().RabbitMQHost,
			config.Get().RabbitMQPort,
		))
	if err != nil {
		log.Fatalf("error while connecting to RabbitMQ: %v\n", err)
	}

	amqpChannel, err := amqpConn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	return &AmqpConnector{
		amqpConn: amqpConn,
		amqpChan: amqpChannel,
	}
}

func (c *AmqpConnector) Close() error {
	if err := c.amqpConn.Close(); err != nil {
		log.Fatalf("Failed to close a connection: %v", err)
		return err
	}

	if err := c.amqpChan.Close(); err != nil {
		log.Fatalf("Failed to close a channel: %v", err)
	}

	return nil
}
