package mongodb

import (
	"context"
	"fmt"
	eventModel "github.com/ferza17/ecommerce-microservices-v2/user-service/model/bson"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *eventMongoRepository) FindEventBySagaIDAndAggregateType(ctx context.Context, sagaID string, aggregateType string) (*eventModel.Event, error) {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "EventMongoDBRepository.FindEventBySagaIDAndAggregateType")
	defer span.End()

	resp := new(eventModel.Event)
	filter := bson.M{"saga_id": sagaID, "aggregate_type": aggregateType}
	if err := r.mongoDB.
		GetCollection(resp.CollectionName()).
		FindOne(ctx, filter).
		Decode(resp); err != nil {
		r.logger.Error(fmt.Sprintf("EventMongoDBRepository.FindEventBySagaIDAndAggregateType : %s", err.Error()))
		span.RecordError(err)
		return nil, err
	}

	return resp, nil
}
