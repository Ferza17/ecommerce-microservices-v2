package rabbitmq

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
	"github.com/google/wire"
	"github.com/rabbitmq/amqp091-go"
)

type (
	IRabbitMQInfrastructure interface {
		SetupQueue(exchange string, topic string, queue string) error

		Consume(ctx context.Context, queue string) (<-chan amqp091.Delivery, error)
		Publish(ctx context.Context, requestId string, exchange string, queue string, message []byte) error
		Close() error
	}
	RabbitMQInfrastructure struct {
		amqpConn                *amqp091.Connection
		channel                 *amqp091.Channel
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  logger.IZapLogger
	}
)

var Set = wire.NewSet(NewRabbitMQInfrastructure)

func NewRabbitMQInfrastructure(
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger) IRabbitMQInfrastructure {
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

	ch, err := amqpConn.Channel()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to open a channel: %v", err))
	}

	return &RabbitMQInfrastructure{
		amqpConn:                amqpConn,
		channel:                 ch,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}

func (c *RabbitMQInfrastructure) Close() error {
	if err := c.amqpConn.Close(); err != nil {
		c.logger.Error(fmt.Sprintf("Failed to close a connection: %v", err))
		return err
	}
	return nil
}
