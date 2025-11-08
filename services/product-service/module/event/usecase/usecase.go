package usecase

import (
	"context"

	kafkaInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/kafka"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	pb "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/event"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IEventUseCase interface {
		AppendEventEnvelope(ctx context.Context, request *pb.EventEnvelope) error
	}

	eventUseCase struct {
		kafkaInfrastructure     kafkaInfrastructure.IKafkaInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  logger.IZapLogger
	}
)

var Set = wire.NewSet(NewEventUseCase)

func NewEventUseCase(
	kafkaInfrastructure kafkaInfrastructure.IKafkaInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
) IEventUseCase {
	return &eventUseCase{
		kafkaInfrastructure:     kafkaInfrastructure,
		telemetryInfrastructure: telemetryInfrastructure,
	}
}
