package rabbitmq

import (
	"context"
	"fmt"
	"log"

	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/config"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/logger"
	"github.com/google/wire"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
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
		logger                  logger.IZapLogger
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
	}
)

var Set = wire.NewSet(NewRabbitMQInfrastructure)

func NewRabbitMQInfrastructure(
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger,
) IRabbitMQInfrastructure {

	amqpConn, err := amqp091.Dial(
		fmt.Sprintf("amqp://%s:%s@%s:%s/",
			config.Get().MessageBrokerRabbitMQ.Username,
			config.Get().MessageBrokerRabbitMQ.Password,
			config.Get().MessageBrokerRabbitMQ.Host,
			config.Get().MessageBrokerRabbitMQ.Port,
		))
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}

	ch, err := amqpConn.Channel()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to open a channel: %s", err))
	}

	if err = ch.Qos(20, 0, false); err != nil {
		log.Fatalf("Failed to set QoS: %s", err)
	}

	c := &RabbitMQInfrastructure{
		amqpConn:                amqpConn,
		channel:                 ch,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
	return c
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
