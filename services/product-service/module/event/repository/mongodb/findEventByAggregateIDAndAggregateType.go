package mongodb

import (
	"context"
	"fmt"

	eventModel "github.com/ferza17/ecommerce-microservices-v2/product-service/model/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *eventMongoRepository) FindEventByAggregateIDAndAggregateType(ctx context.Context, aggregateID string, aggregateType string) (*eventModel.Event, error) {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "EventMongoDBRepository.FindEventByAggregateIDAndAggregateType")
	defer span.End()

	resp := new(eventModel.Event)

	filter := bson.M{"aggregate_id": aggregateID, "aggregate_type": aggregateType}
	opts := options.FindOne().SetSort(bson.D{{"timestamp", -1}})
	if err := r.mongoDB.
		GetCollection(resp.CollectionName()).
		FindOne(ctx, filter, opts).
		Decode(resp); err != nil {
		r.logger.Error(fmt.Sprintf("EventMongoDBRepository.FindEventByAggregateIDAndAggregateType : %s", err.Error()))
		span.RecordError(err)
		return nil, err
	}
	return resp, nil
}
