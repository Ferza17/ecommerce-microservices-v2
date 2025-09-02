package presenter

import (
	"context"

	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/telemetry"
	pb "github.com/ferza17/ecommerce-microservices-v2/event-store-service/model/rpc/gen/v1/event"
	eventUseCase "github.com/ferza17/ecommerce-microservices-v2/event-store-service/module/event/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/logger"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

type (
	IEventPresenter interface {
		Append(ctx context.Context, req *pb.AppendRequest) (*pb.AppendResponse, error)
		ReadByAggregate(ctx context.Context, req *pb.ReadByAggregateRequest) (*pb.ReadByAggregateResponse, error)
		ReadByType(ctx context.Context, req *pb.ReadByTypeRequest) (*pb.ReadByTypeResponse, error)
		// Server-streaming: continuous feed for projectors
		Subscribe(req *pb.SubscribeRequest, srv grpc.ServerStreamingServer[pb.Event]) error
		GetSnapshot(ctx context.Context, req *pb.GetSnapshotRequest) (*pb.GetSnapshotResponse, error)
		PutSnapshot(ctx context.Context, req *pb.PutSnapshotRequest) (*pb.PutSnapshotResponse, error)
	}
	eventPresenter struct {
		pb.UnimplementedEventStoreServer
		eventUseCase eventUseCase.IEventUseCase
		telemetry    telemetry.ITelemetryInfrastructure
		logger       logger.IZapLogger
	}
)

var Set = wire.NewSet(NewEventPresenter)

func NewEventPresenter(
	eventUseCase eventUseCase.IEventUseCase,
	telemetry telemetry.ITelemetryInfrastructure,
	logger logger.IZapLogger,
) IEventPresenter {
	return &eventPresenter{
		eventUseCase: eventUseCase,
		telemetry:    telemetry,
		logger:       logger,
	}
}
