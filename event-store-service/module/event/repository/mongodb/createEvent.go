package mongodb

import (
	"context"
	"fmt"

	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/model/bson"
)

func (r *eventRepository) CreateEvent(ctx context.Context, requestId string, request *bson.Event) (*bson.Event, error) {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "EventRepository.CreateEvent")
	defer span.End()

	if _, err := r.collection.
		InsertOne(ctx, request); err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , Failed to insert event: %s", requestId, err.Error()))
		return nil, err
	}

	return request, nil
}
