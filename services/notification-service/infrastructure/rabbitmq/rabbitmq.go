package rabbitmq

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/config"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	"github.com/google/wire"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"log"
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
		logger                  logger.IZapLogger
		channel                 *amqp091.Channel
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
	}
)

var Set = wire.NewSet(NewRabbitMQInfrastructure)

func NewRabbitMQInfrastructure(
	logger logger.IZapLogger,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
) IRabbitMQInfrastructure {
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

	if err = ch.Qos(20, 0, false); err != nil {
		log.Fatalf("Failed to set QoS: %s", err)
	}

	return &RabbitMQInfrastructure{
		amqpConn:                amqpConn,
		channel:                 ch,
		logger:                  logger,
		telemetryInfrastructure: telemetryInfrastructure,
	}
}

func (c *RabbitMQInfrastructure) Close() error {
	if err := c.channel.Close(); err != nil {
		c.logger.Error("Failed to close the channel", zap.Error(err))
		return err
	}

	if err := c.amqpConn.Close(); err != nil {
		c.logger.Error(fmt.Sprintf("Failed to close a connection: %v", err))
		return err
	}

	return nil
}
