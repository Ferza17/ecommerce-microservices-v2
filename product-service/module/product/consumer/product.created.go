package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/enum"
	pb "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/context"
	pkgMetric "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/metric"
	"github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"sync"
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

	deliveries, err := amqpChannel.Consume(
		config.Get().QueueProductCreated,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		c.logger.Error(fmt.Sprintf("failed to consume : %v", zap.Error(err)))
		return err
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(deliveries <-chan amqp091.Delivery) {
		defer wg.Done()
	messages:
		for d := range deliveries {
			var (
				request           pb.UpdateProductByIdRequest
				requestId         string
				newCtx, cancelCtx = context.WithTimeout(ctx, 20)
			)
			for key, value := range d.Headers {
				if key == pkgContext.CtxKeyRequestID {
					requestId = value.(string)
					newCtx = pkgContext.SetRequestIDToContext(ctx, requestId)
				}

				if key == pkgContext.CtxKeyAuthorization {
					newCtx = pkgContext.SetTokenAuthorizationToContext(ctx, value.(string))
				}

			}

			newCtx, span := c.telemetryInfrastructure.StartSpanFromRabbitMQHeader(newCtx, d.Headers, "ProductConsumer.ProductCreated")
			span.SetAttributes(attribute.String("messaging.destination", config.Get().QueueProductCreated))
			span.SetAttributes(attribute.String(pkgContext.CtxKeyRequestID, requestId))

			switch d.ContentType {
			case enum.XProtobuf.String():
				if err = proto.Unmarshal(d.Body, &request); err != nil {
					c.logger.Error(fmt.Sprintf("requsetID : %s , failed to unmarshal request : %v", requestId, zap.Error(err)))
					pkgMetric.RabbitmqMessagesConsumed.WithLabelValues(config.Get().QueueProductCreated, "failed").Inc()
					span.RecordError(err)
					span.End()
					d.Nack(false, true)
					cancelCtx()
					continue messages
				}
			case enum.JSON.String():
				if err = json.Unmarshal(d.Body, &request); err != nil {
					c.logger.Error(fmt.Sprintf("failed to unmarshal request : %v", zap.Error(err)))
					pkgMetric.RabbitmqMessagesConsumed.WithLabelValues(config.Get().QueueProductCreated, "failed").Inc()
					span.RecordError(err)
					span.End()
					d.Nack(false, true)
					cancelCtx()
					continue messages
				}
			default:
				c.logger.Error(fmt.Sprintf("failed to get request id"))
				pkgMetric.RabbitmqMessagesConsumed.WithLabelValues(config.Get().QueueProductCreated, "failed").Inc()
				span.RecordError(err)
				span.End()
				d.Nack(false, true)
				cancelCtx()
				continue messages
			}

			c.logger.Info(fmt.Sprintf("received a %s message: %s", d.RoutingKey, d.Body))
			if _, err = c.productUseCase.UpdateProductById(newCtx, requestId, &request); err != nil {
				pkgMetric.RabbitmqMessagesConsumed.WithLabelValues(config.Get().QueueProductCreated, "failed").Inc()
				span.RecordError(err)
				span.End()
				d.Nack(false, true)
				cancelCtx()
				continue messages
			}

			pkgMetric.RabbitmqMessagesConsumed.WithLabelValues(config.Get().QueueProductCreated, "success").Inc()
			span.End()
			cancelCtx()
		}
	}(deliveries)

	wg.Wait()
	return nil
}
