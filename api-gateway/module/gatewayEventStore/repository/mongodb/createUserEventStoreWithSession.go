package mongodb

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/enum"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *gatewayEventStoreRepository) CreateGatewayEventStoreWithSession(ctx context.Context, requestId string, event *bson.Event, session mongo.Session) (string, error) {
	var (
		resp string
	)

	if err := mongo.WithSession(ctx, session, func(sessionContext mongo.SessionContext) error {
		gatewayEventStoreCollection := r.MongoDB.GetCollection(enum.DatabaseEventStore, enum.CollectionUserEvent)
		result, err := gatewayEventStoreCollection.InsertOne(ctx, event)
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
