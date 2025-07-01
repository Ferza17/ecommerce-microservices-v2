package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/enum"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	"github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

func (c *productConsumer) ProductCreated(ctx context.Context) error {
	amqpChannel, err := c.rabbitMQInfrastructure.GetConnection().Channel()
	if err != nil {
		c.logger.Error("ProductConsumer.ProductCreated", zap.Error(err))
		return err
	}

	if err = amqpChannel.ExchangeDeclare(
		config.Get().ExchangeProduct,
		amqp091.ExchangeDirect,
		true,
		false,
		false,
		true,
		nil,
	); err != nil {
		c.logger.Error("ProductConsumer.ProductCreated", zap.String("Exchange", config.Get().ExchangeProduct), zap.Error(err))
		return err
	}

	if err = amqpChannel.QueueBind(
		config.Get().QueueProductCreated,
		config.Get().QueueProductCreated,
		config.Get().ExchangeProduct,
		false,
		nil,
	); err != nil {
		c.logger.Error("ProductConsumer.ProductCreated", zap.String("Exchange", config.Get().ExchangeProduct), zap.String("Queue", config.Get().QueueProductCreated), zap.Error(err))
		return err
	}

	msgs, err := amqpChannel.Consume(
		config.Get().QueueProductCreated,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
	messages:
		for d := range msgs {
			var (
				request   productRpc.CreateProductRequest
				requestId string
			)
			carrier := propagation.MapCarrier{}
			for key, value := range d.Headers {
				if key == enum.XRequestIDHeader.String() {
					requestId = value.(string)
				}

				if strVal, ok := value.(string); ok {
					carrier[key] = strVal
				}
			}
			ctx = otel.GetTextMapPropagator().Extract(context.Background(), carrier)
			ctx, span := c.telemetryInfrastructure.Tracer(ctx, "Consumer.ProductCreated")

			switch d.ContentType {
			case enum.XProtobuf.String():
				if err = proto.Unmarshal(d.Body, &request); err != nil {
					c.logger.Error(fmt.Sprintf("requsetID : %s , failed to unmarshal request : %v", requestId, zap.Error(err)))
					span.End()
					continue messages
				}
			case enum.JSON.String():
				if err = json.Unmarshal(d.Body, &request); err != nil {
					c.logger.Error(fmt.Sprintf("failed to unmarshal request : %v", zap.Error(err)))
					span.End()
					continue messages
				}
			default:
				c.logger.Error(fmt.Sprintf("failed to get ContentType"))
				span.End()
				continue messages
			}

			if _, err = c.productUseCase.CreateProduct(ctx, requestId, &request); err != nil {
				c.logger.Error(fmt.Sprintf("failed to create user : %v", zap.Error(err)))
				span.End()
				continue messages
			}
			span.End()
		}
	}()

	<-ctx.Done()

	return nil
}
