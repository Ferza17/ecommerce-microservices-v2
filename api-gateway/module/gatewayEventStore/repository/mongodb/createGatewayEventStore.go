package mongodb

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/enum"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *gatewayEventStoreRepository) CreateGatewayEventStore(ctx context.Context, requestId string, req *bson.Event) (string, error) {
	gatewayEventStoreCollection := r.MongoDB.GetCollection(enum.DatabaseEventStore, enum.CollectionUserEvent)

	req.ID = primitive.NewObjectID()
	if _, err := gatewayEventStoreCollection.InsertOne(ctx, req); err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error creating userEventStore event: %v", requestId, err))
		return "", nil
	}

	return req.ID.String(), nil
}
