package mongodb

import (
	"context"
	"fmt"

	eventModel "github.com/ferza17/ecommerce-microservices-v2/product-service/model/bson"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *eventMongoRepository) FindEventBySagaID(ctx context.Context, sagaID string) (*eventModel.Event, error) {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "EventMongoDBRepository.FindEventBySagaID")
	defer span.End()

	resp := new(eventModel.Event)
	if err := r.mongoDB.
		GetCollection(resp.CollectionName()).
		FindOne(ctx, bson.M{"saga_id": sagaID}).
		Decode(resp); err != nil {
		r.logger.Error(fmt.Sprintf("EventMongoDBRepository.FindEventBySagaID : %s", err.Error()))
		span.RecordError(err)
		return nil, err
	}

	return resp, nil
}
