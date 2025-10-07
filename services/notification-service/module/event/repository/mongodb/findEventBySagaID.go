package mongodb

import (
	"context"
	"fmt"

	"github.com/ferza17/ecommerce-microservices-v2/notification-service/enum"
	eventModel "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *eventMongoRepository) FindEventBySagaID(ctx context.Context, sagaID string) (*eventModel.Event, error) {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "EventMongoDBRepository.FindEventBySagaID")
	defer span.End()

	resp := new(eventModel.Event)
	filter := bson.M{"saga_id": sagaID}
	opts := options.FindOne().SetSort(bson.D{{"timestamp", -1}})
	if err := r.mongoDB.
		GetCollection(enum.DatabaseNotification, resp.CollectionName()). // TODO: Change to database notification_event_stores
		FindOne(ctx, filter, opts).
		Decode(resp); err != nil {
		r.logger.Error(fmt.Sprintf("EventMongoDBRepository.FindEventBySagaID : %s", err.Error()))
		span.RecordError(err)
		return nil, err
	}

	return resp, nil
}
