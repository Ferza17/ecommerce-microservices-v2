package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/enum"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/context"
	pkgMetric "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/metric"
	"github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"sync"
)

func (c *paymentConsumer) PaymentOrderDelayedCancelled(ctx context.Context) error {
	amqpChannel, err := c.rabbitmq.GetConnection().Channel()
	if err != nil {
		c.logger.Error(fmt.Sprintf("Failed to create a channel: %v", err))
		return err
	}

	if err = amqpChannel.ExchangeDeclare(
		config.Get().ExchangePaymentDelayed,
		"x-delayed-message",
		true,
		false,
		false,
		true,
		amqp091.Table{
			enum.XDelayedType.String(): "direct",
		},
	); err != nil {
		c.logger.Error(fmt.Sprintf("failed to declare exchange : %v", zap.Error(err)))
		return err
	}

	if err = amqpChannel.QueueBind(
		config.Get().QueuePaymentOrderDelayedCancelled,
		"",
		config.Get().ExchangePaymentDelayed,
		false,
		nil,
	); err != nil {
		c.logger.Error(fmt.Sprintf("failed to bind queue : %v", zap.Error(err)))
		return err
	}

	deliveries, err := amqpChannel.Consume(
		config.Get().QueuePaymentOrderDelayedCancelled,
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
				request           paymentRpc.PaymentOrderDelayedCancelledRequest
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

			newCtx, span := c.telemetryInfrastructure.StartSpanFromRabbitMQHeader(newCtx, d.Headers, "PaymentConsumer.PaymentOrderDelayedCancelled")
			span.SetAttributes(attribute.String("messaging.destination", config.Get().QueuePaymentOrderDelayedCancelled))
			span.SetAttributes(attribute.String(pkgContext.CtxKeyRequestID, requestId))

			switch d.ContentType {
			case enum.XProtobuf.String():
				if err = proto.Unmarshal(d.Body, &request); err != nil {
					c.logger.Error(fmt.Sprintf("requsetID : %s , failed to unmarshal request : %v", requestId, zap.Error(err)))
					pkgMetric.RabbitmqMessagesConsumed.WithLabelValues(config.Get().QueuePaymentOrderDelayedCancelled, "failed").Inc()
					span.RecordError(err)
					span.End()
					d.Nack(false, true)
					cancelCtx()
					continue messages
				}
			case enum.JSON.String():
				if err = json.Unmarshal(d.Body, &request); err != nil {
					c.logger.Error(fmt.Sprintf("failed to unmarshal request : %v", zap.Error(err)))
					pkgMetric.RabbitmqMessagesConsumed.WithLabelValues(config.Get().QueuePaymentOrderDelayedCancelled, "failed").Inc()
					span.RecordError(err)
					span.End()
					d.Nack(false, true)
					cancelCtx()
					continue messages
				}
			default:
				c.logger.Error(fmt.Sprintf("failed to get request id"))
				pkgMetric.RabbitmqMessagesConsumed.WithLabelValues(config.Get().QueuePaymentOrderDelayedCancelled, "failed").Inc()
				span.RecordError(err)
				span.End()
				d.Nack(false, true)
				cancelCtx()
				continue messages
			}

			c.logger.Info(fmt.Sprintf("received a %s message: %s", d.RoutingKey, d.Body))
			if err = c.paymentUseCase.PaymentOrderDelayedCancelled(newCtx, requestId, &request); err != nil {
				pkgMetric.RabbitmqMessagesConsumed.WithLabelValues(config.Get().QueuePaymentOrderDelayedCancelled, "failed").Inc()
				span.RecordError(err)
				span.End()
				d.Nack(false, true)
				cancelCtx()
				continue messages
			}

			pkgMetric.RabbitmqMessagesConsumed.WithLabelValues(config.Get().QueuePaymentOrderDelayedCancelled, "success").Inc()
			span.End()
			cancelCtx()
		}
	}(deliveries)

	wg.Wait()
	return nil
}
