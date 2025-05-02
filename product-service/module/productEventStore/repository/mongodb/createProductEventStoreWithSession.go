package mongodb

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *productEventStoreRepository) CreateProductEventStoreWithSession(ctx context.Context, requestId string, req *bson.Event, session mongo.Session) (string, error) {
	var (
		resp string
	)

	if err := mongo.WithSession(ctx, session, func(sessionContext mongo.SessionContext) error {
		productEventStoreCollection := r.connector.MongoClient.Database(databaseEventStore).Collection(collectionProductEvent)
		result, err := productEventStoreCollection.InsertOne(ctx, req)
		if err != nil {
			r.logger.Error(fmt.Sprintf("requestId : %s , error creating ProductEventStore event: %v", requestId, err))
			return err
		}
		resp = result.InsertedID.(string)
		return nil
	}); err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error creating ProductEventStore event: %v", requestId, err))
	}

	return resp, nil
}
