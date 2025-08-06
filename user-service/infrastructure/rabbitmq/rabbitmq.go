package rabbitmq

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/temporal"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
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
		channel                 *amqp091.Channel
		logger                  logger.IZapLogger
		temporal                temporal.ITemporalInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
	}
)

var Set = wire.NewSet(NewRabbitMQInfrastructure)

func NewRabbitMQInfrastructure(
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	temporal temporal.ITemporalInfrastructure,
	logger logger.IZapLogger,
) IRabbitMQInfrastructure {

	amqpConn, err := amqp091.Dial(
		fmt.Sprintf("amqp://%s:%s@%s:%s/",
			config.Get().RabbitMQUsername,
			config.Get().RabbitMQPassword,
			config.Get().RabbitMQHost,
			config.Get().RabbitMQPort,
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
		temporal:                temporal,
	}
	c.temporal = c.temporal.RegisterActivity(c.Publish)
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
