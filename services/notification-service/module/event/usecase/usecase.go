package usecase

import (
	"context"

	kafkaInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/kafka"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/telemetry"
	pb "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/v1/event"
	eventMongoDBRepository "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/event/repository/mongodb"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IEventUseCase interface {
		AppendEvent(ctx context.Context, request *pb.Event) error
	}

	eventUseCase struct {
		eventMongoDBRepository  eventMongoDBRepository.IEventMongoRepository
		kafkaInfrastructure     kafkaInfrastructure.IKafkaInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  logger.IZapLogger
	}
)

var Set = wire.NewSet(NewEventUseCase)

func NewEventUseCase(
	eventMongoDBRepository eventMongoDBRepository.IEventMongoRepository,
	kafkaInfrastructure kafkaInfrastructure.IKafkaInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
) IEventUseCase {
	return &eventUseCase{
		eventMongoDBRepository:  eventMongoDBRepository,
		kafkaInfrastructure:     kafkaInfrastructure,
		telemetryInfrastructure: telemetryInfrastructure,
	}
}
