package rabbitmq

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/config"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/enum"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
	"github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
	"time"
)

type (
	IRabbitMQInfrastructure interface {
		Publish(ctx context.Context, requestId string, exchange string, queue string, message []byte) error
		Close() error
	}
	RabbitMQInfrastructure struct {
		amqpConn                *amqp091.Connection
		logger                  pkg.IZapLogger
		cb                      pkg.ICircuitBreaker
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
	}
)

func NewRabbitMQInfrastructure(
	logger pkg.IZapLogger,
	//cb pkg.ICircuitBreaker,
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

	return &RabbitMQInfrastructure{
		amqpConn: amqpConn,
		logger:   logger,
		//cb:                      cb,
		telemetryInfrastructure: telemetryInfrastructure,
	}
}

func (c *RabbitMQInfrastructure) Close() error {
	if err := c.amqpConn.Close(); err != nil {
		c.logger.Error(fmt.Sprintf("Failed to close a connection: %v", err))
		return err
	}
	return nil
}

func (c *RabbitMQInfrastructure) Publish(ctx context.Context, requestId string, exchange string, queue string, message []byte) error {
	ctx, span := c.telemetryInfrastructure.Tracer(ctx, "RabbitMQInfrastructure.Publish")

	amqpChannel, err := c.amqpConn.Channel()
	if err != nil {
		c.logger.Error(fmt.Sprintf("Failed to create a channel: %v", err))
		return err
	}
	defer func(amqpChannel *amqp091.Channel) {
		span.AddEvent(queue)
		defer span.End()
		if err = amqpChannel.Close(); err != nil {
			c.logger.Error(fmt.Sprintf("Failed to close a channel: %v", err))
		}
	}(amqpChannel)

	if err = amqpChannel.ExchangeDeclare(
		exchange,
		amqp091.ExchangeDirect,
		true,
		false,
		false,
		true,
		nil,
	); err != nil {
		c.logger.Error(fmt.Sprintf("failed to declare exchange : %v", zap.Error(err)))
		return err
	}

	carrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(ctx, carrier)
	headers := amqp091.Table{}
	for k, v := range carrier {
		headers[k] = v
	}
	headers[enum.XRequestIDHeader.String()] = requestId

	// Publish message
	if err = amqpChannel.PublishWithContext(
		ctx,
		exchange,
		queue,
		false,
		false,
		amqp091.Publishing{
			ContentType:  enum.XProtobuf.String(),
			DeliveryMode: amqp091.Transient,
			Timestamp:    time.Now(),
			Body:         message,
			Headers:      headers,
		},
	); err != nil {
		c.logger.Error(fmt.Sprintf("Failed to publish a message: %v", err))
		return err
	}

	return nil
}
