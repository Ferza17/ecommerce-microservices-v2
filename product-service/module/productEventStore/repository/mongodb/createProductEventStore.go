package mongodb

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *productEventStoreRepository) CreateProductEventStore(ctx context.Context, requestId string, req *bson.Event) (string, error) {
	productEventStoreCollection := r.connector.MongoClient.Database(databaseEventStore).Collection(collectionProductEvent)

	req.ID = primitive.NewObjectID()
	if _, err := productEventStoreCollection.InsertOne(ctx, req); err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error creating productEventStore event: %v", requestId, err))
		return "", nil
	}

	return req.ID.String(), nil
}
