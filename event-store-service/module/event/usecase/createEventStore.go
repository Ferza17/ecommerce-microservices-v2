package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/model/bson"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/model/rpc/pb"
	"time"
)

func (u *eventUseCase) CreateEventStore(ctx context.Context, requestId string, req *pb.EventStore) (*pb.CreateEventStoreResponse, error) {
	var (
		now       = time.Now().UTC()
		bsonEvent = &bson.Event{
			SagaID:    requestId,
			Entity:    req.Entity,
			EntityID:  req.EntityId,
			EventType: req.EventType,
			Status:    req.Status,
			CreatedAt: now,
			UpdatedAt: now,
		}
	)

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

	return &pb.CreateEventStoreResponse{
		Id: result,
	}, err
}
