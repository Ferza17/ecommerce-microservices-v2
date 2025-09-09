package usecase

import (
	"context"

	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/telemetry"
	pb "github.com/ferza17/ecommerce-microservices-v2/event-store-service/model/rpc/gen/v1/event"
	eventRepository "github.com/ferza17/ecommerce-microservices-v2/event-store-service/module/event/repository/mongodb"
	pkgLogger "github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IEventUseCase interface {
		Append(ctx context.Context, requestId string, req *pb.AppendRequest) (*pb.AppendResponse, error)
		ReadByAggregate(ctx context.Context, requestId string, req *pb.ReadByAggregateRequest) (*pb.ReadByAggregateResponse, error)
		ReadByType(ctx context.Context, requestId string, req *pb.ReadByTypeRequest) (*pb.ReadByTypeResponse, error)

		GetSnapshot(ctx context.Context, requestId string, req *pb.GetSnapshotRequest) (*pb.GetSnapshotResponse, error)
		PutSnapshot(ctx context.Context, requestId string, req *pb.PutSnapshotRequest) (*pb.PutSnapshotResponse, error)
	}

	eventUseCase struct {
		eventRepository         eventRepository.IEventRepository
		rabbitMQInfrastructure  rabbitmqInfrastructure.IRabbitMQInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  pkgLogger.IZapLogger
	}
)

var Set = wire.NewSet(NewEventStoreUseCase)

func NewEventStoreUseCase(
	eventStoreRepository eventRepository.IEventRepository,
	rabbitMQInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkgLogger.IZapLogger,
) IEventUseCase {
	return &eventUseCase{
		eventRepository:         eventStoreRepository,
		rabbitMQInfrastructure:  rabbitMQInfrastructure,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
