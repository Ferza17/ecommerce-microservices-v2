package consumer

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	kafkaInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/kafka"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	productUseCase "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IProductConsumer interface {
		SnapshotProductsProductCreated(ctx context.Context, message *kafka.Message) error
		SnapshotProductsProductUpdated(ctx context.Context, message *kafka.Message) error
		SnapshotProductsProductDeleted(ctx context.Context, message *kafka.Message) error
	}

	productConsumer struct {
		kafkaInfrastructure     kafkaInfrastructure.IKafkaInfrastructure
		productUseCase          productUseCase.IProductUseCase
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  logger.IZapLogger
	}
)

var Set = wire.NewSet(NewProductConsumer)

func NewProductConsumer(
	kafkaInfrastructure kafkaInfrastructure.IKafkaInfrastructure,
	productUseCase productUseCase.IProductUseCase,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger,
) IProductConsumer {
	return &productConsumer{
		kafkaInfrastructure:     kafkaInfrastructure,
		productUseCase:          productUseCase,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
