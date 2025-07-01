package rabbitmq

import (
	"context"
	"fmt"
	productConsumer "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/consumer"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
	"github.com/google/wire"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"sync"
)

type (
	RabbitMQTransport struct {
		amqpConn        *amqp091.Connection
		logger          logger.IZapLogger
		productConsumer productConsumer.IProductConsumer
	}
)

var Set = wire.NewSet(NewServer)

func NewServer(
	logger logger.IZapLogger,
	productConsumer productConsumer.IProductConsumer,
) *RabbitMQTransport {
	return &RabbitMQTransport{
		logger:          logger,
		productConsumer: productConsumer,
	}
}

func (srv *RabbitMQTransport) Serve(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := srv.productConsumer.ProductCreated(ctx); err != nil {
			srv.logger.Error(fmt.Sprintf("failed to ProductCreated : %s", zap.Error(err).String))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := srv.productConsumer.ProductUpdated(ctx); err != nil {
			srv.logger.Error(fmt.Sprintf("failed to ProductUpdated : %s", zap.Error(err).String))
		}
	}()

	wg.Wait()
	cancel()
}
