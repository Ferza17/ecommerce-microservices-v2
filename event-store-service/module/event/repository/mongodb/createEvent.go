package mongodb

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/enum"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/model/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *eventRepository) CreateEvent(ctx context.Context, requestId string, req *bson.Event) (string, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "Repository.CreateEvent")
	defer span.End()
	req.ID = primitive.NewObjectID()
	if _, err := r.mongoDB.
		GetCollection(enum.DatabaseEventStore, enum.CollectionEvent).
		InsertOne(ctx, req); err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s ,Failed to insert event: %s", requestId, err.Error()))
		return "", err
	}
	return req.ID.String(), nil
}
