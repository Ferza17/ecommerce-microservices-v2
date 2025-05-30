package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/model/bson"
	eventRpc "github.com/ferza17/ecommerce-microservices-v2/event-store-service/model/rpc/gen/event/v1"
	"time"
)

func (u *eventUseCase) CreateEventStore(ctx context.Context, requestId string, req *eventRpc.EventStore) (*eventRpc.CreateEventStoreResponse, error) {
	var (
		now       = time.Now().UTC()
		bsonEvent = &bson.Event{
			RequestID: requestId,
			Service:   req.Service,
			EventType: req.EventType,
			Status:    req.Status,
			CreatedAt: now,
			UpdatedAt: now,
		}
	)
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.CreateEventStore")
	defer span.End()

	if req.Payload != nil {
		p := req.Payload.AsMap()
		bsonEvent.Payload = &p
	}

	if req.PreviousState != nil {
		p := req.PreviousState.AsMap()
		bsonEvent.PreviousState = &p
	}

	result, err := u.eventRepository.CreateEvent(ctx, requestId, bsonEvent)
	if err != nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error creating userEventStore event: %v", requestId, err))
	}

	return &eventRpc.CreateEventStoreResponse{
		Id: result,
	}, err
}
