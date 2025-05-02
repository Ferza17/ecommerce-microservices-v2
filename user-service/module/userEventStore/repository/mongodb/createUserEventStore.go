package mongodb

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *userEventStoreRepository) CreateUserEventStore(ctx context.Context, requestId string, req *bson.Event) (string, error) {
	userEventStoreCollection := r.connector.MongoClient.Database(databaseEventStore).Collection(collectionUserEvent)

	req.ID = primitive.NewObjectID()
	if _, err := userEventStoreCollection.InsertOne(ctx, req); err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error creating userEventStore event: %v", requestId, err))
		return "", nil
	}

	return req.ID.String(), nil
}
