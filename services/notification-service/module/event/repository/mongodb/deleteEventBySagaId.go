package mongodb

import (
	"context"
	"fmt"

	"github.com/ferza17/ecommerce-microservices-v2/notification-service/enum"
	eventModel "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/bson"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *eventMongoRepository) DeleteEventBySagaId(ctx context.Context, sagaID string) error {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "EventMongoDBRepository.FindEventBySagaIDAndAggregateType")
	defer span.End()

	filter := bson.M{"saga_id": sagaID}
	if _, err := r.mongoDB.
		GetCollection(enum.DatabaseNotification, eventModel.Event{}.CollectionName()). // TODO: Change to DB notification_event_stores
		DeleteOne(ctx, filter); err != nil {
		r.logger.Error(fmt.Sprintf("EventMongoDBRepository.FindEventByAggregateIDAndAggregateType : %s", err.Error()))
		span.RecordError(err)
		return err
	}

	return nil
}
