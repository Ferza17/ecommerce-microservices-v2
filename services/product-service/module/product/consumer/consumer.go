package consumer

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	productUseCase "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
	"github.com/google/wire"
	"github.com/rabbitmq/amqp091-go"
)

type (
	IProductConsumer interface {
		ProductCreated(ctx context.Context, d *amqp091.Delivery) error
		ProductUpdated(ctx context.Context, d *amqp091.Delivery) error
		ProductDeleted(ctx context.Context, d *amqp091.Delivery) error
	}

	productConsumer struct {
		rabbitMQInfrastructure  rabbitmqInfrastructure.IRabbitMQInfrastructure
		productUseCase          productUseCase.IProductUseCase
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  logger.IZapLogger
	}
)

var Set = wire.NewSet(NewProductConsumer)

func NewProductConsumer(
	rabbitMQInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	productUseCase productUseCase.IProductUseCase,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger,
) IProductConsumer {
	return &productConsumer{
		rabbitMQInfrastructure:  rabbitMQInfrastructure,
		productUseCase:          productUseCase,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
