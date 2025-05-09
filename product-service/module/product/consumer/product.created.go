package consumer

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/enum"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

func (c *productConsumer) ProductCreated(ctx context.Context) error {
	var (
		requestId string
		ok        bool
	)

	if err := c.amqpChannel.ExchangeDeclare(
		enum.ProductExchange.String(),
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

	if err := c.amqpChannel.QueueBind(
		enum.QueueProduct.String(),
		enum.PRODUCT_CREATED.String(),
		enum.ProductExchange.String(),
		false,
		nil,
	); err != nil {
		c.logger.Error(fmt.Sprintf("failed to bind queue : %v", zap.Error(err)))
		return err
	}

	msgs, err := c.amqpChannel.Consume(
		enum.QueueProduct.String(),
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
			var request pb.CreateProductRequest
			if requestId, ok = d.Headers[enum.XRequestIDHeader.String()].(string); !ok {
				c.logger.Error("failed to get request id")
				continue messages
			}

			if err = proto.Unmarshal(d.Body, &request); err != nil {
				c.logger.Error(fmt.Sprintf("requsetID : %s , failed to unmarshal request : %v", requestId, zap.Error(err)))
				continue messages
			}

			if _, err = c.productUseCase.CreateProduct(ctx, requestId, &request); err != nil {
				c.logger.Error(fmt.Sprintf("failed to create user : %v", zap.Error(err)))
				continue messages
			}
		}
	}()

	<-ctx.Done()

	return nil
}
