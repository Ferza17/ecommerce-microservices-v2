package usecase

import (
	"context"

	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/model/bson"
	pb "github.com/ferza17/ecommerce-microservices-v2/event-store-service/model/rpc/gen/v1/event"
)

func (u *eventUseCase) Append(ctx context.Context, requestId string, req *pb.AppendRequest) (*pb.AppendResponse, error) {
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "EventRepository.Append")
	defer span.End()

	for _, event := range req.Events {
		if _, err := u.eventRepository.CreateEvent(ctx, requestId, bson.EventFromProto(event)); err != nil {
			span.RecordError(err)
			return nil, err
		}
	}

	return &pb.AppendResponse{
		NextVersion: req.ExpectedVersion + 1,
	}, nil
}
