package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/enum"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	"github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

func (c *paymentConsumer) PaymentOrderCreated(ctx context.Context) error {
	amqpChannel, err := c.rabbitmq.GetConnection().Channel()
	if err != nil {
		c.logger.Error(fmt.Sprintf("Failed to create a channel: %v", err))
		return err
	}

	if err = amqpChannel.ExchangeDeclare(
		config.Get().ExchangePaymentDirect,
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

	if err = amqpChannel.QueueBind(
		config.Get().QueuePaymentOrderCreated,
		config.Get().QueuePaymentOrderCreated,
		config.Get().ExchangePaymentDirect,
		false,
		nil,
	); err != nil {
		c.logger.Error(fmt.Sprintf("failed to bind queue : %v", zap.Error(err)))
		return err
	}

	msgs, err := amqpChannel.Consume(
		config.Get().QueuePaymentOrderCreated,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	ctx, cancel := context.WithCancel(context.Background())
	go func(deliveries <-chan amqp091.Delivery) {
		defer cancel()
	messages:
		for d := range deliveries {
			var (
				request   paymentRpc.CreatePaymentRequest
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
			ctx := otel.GetTextMapPropagator().Extract(context.Background(), carrier)
			ctx, span := c.telemetryInfrastructure.Tracer(ctx, "AuthConsumer.UserLogin")

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
				c.logger.Error(fmt.Sprintf("failed to get request id"))
				span.End()
				continue messages
			}

			if err = c.paymentUseCase.CreatePayment(ctx, requestId, &request); err != nil {
				c.logger.Error(fmt.Sprintf("failed to CreatePayment : %v", zap.Error(err)))
				span.End()
				continue messages
			}
			span.End()
		}
	}(msgs)

	<-ctx.Done()
	return nil
}
