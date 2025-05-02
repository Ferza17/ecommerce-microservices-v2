package mongodb

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *userEventStoreRepository) CreateUserEventStoreWithSession(ctx context.Context, requestId string, event *bson.Event, session mongo.Session) (string, error) {
	var (
		resp string
	)

	if err := mongo.WithSession(ctx, session, func(sessionContext mongo.SessionContext) error {
		userEventStoreCollection := r.connector.MongoClient.Database(databaseEventStore).Collection(collectionUserEvent)
		result, err := userEventStoreCollection.InsertOne(ctx, event)
		if err != nil {
			r.logger.Error(fmt.Sprintf("requestId : %s , error creating userEventStore event: %v", requestId, err))
			return err
		}
		resp = result.InsertedID.(string)
		return nil
	}); err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error creating userEventStore event: %v", requestId, err))
	}

	return resp, nil
}
