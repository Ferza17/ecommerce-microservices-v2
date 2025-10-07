package mongodb

import (
	"context"
	"fmt"

	eventModel "github.com/ferza17/ecommerce-microservices-v2/product-service/model/bson"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *eventMongoRepository) FindEventsBySagaID(ctx context.Context, sagaID string) ([]*eventModel.Event, error) {
	var (
		err    error
		events []*eventModel.Event
	)

	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "EventMongoDBRepository.FindEventBySagaID")
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()

	cursor, err := r.mongoDB.GetCollection(eventModel.Event{}.CollectionName()).Find(ctx, bson.M{"saga_id": sagaID})
	if err != nil {
		// handle canceled context cleanly
		if ctx.Err() != nil {
			return nil, fmt.Errorf("context canceled while querying MongoDB: %w", ctx.Err())
		}
		return nil, fmt.Errorf("failed to find events by saga_id: %w", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var e eventModel.Event
		if err = cursor.Decode(&e); err != nil {
			return nil, fmt.Errorf("failed to decode event document: %w", err)
		}
		events = append(events, &e)
	}

	if err = cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return events, nil
}
