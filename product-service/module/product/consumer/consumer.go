package consumer

import (
	"context"
	productUseCase "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg"
	"github.com/rabbitmq/amqp091-go"
)

type (
	IProductConsumer interface {
		ProductCreated(ctx context.Context) error
		ProductUpdated(ctx context.Context) error
		ProductDeleted(ctx context.Context) error
	}

	productConsumer struct {
		amqpChannel    *amqp091.Channel
		productUseCase productUseCase.IProductUseCase
		logger         pkg.IZapLogger
	}
)

func NewProductConsumer(amqpChannel *amqp091.Channel, productUseCase productUseCase.IProductUseCase, logger pkg.IZapLogger) IProductConsumer {
	return &productConsumer{
		amqpChannel:    amqpChannel,
		productUseCase: productUseCase,
		logger:         logger,
	}
}
