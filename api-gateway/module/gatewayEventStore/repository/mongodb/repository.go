package mongodb

import (
	"context"
	mongodbInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/mongodb"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/bson"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	IGatewayEventStoreRepository interface {
		CreateGatewayEventStore(ctx context.Context, requestId string, req *bson.Event) (string, error)
		CreateGatewayEventStoreWithSession(ctx context.Context, requestId string, event *bson.Event, session mongo.Session) (string, error)
	}

	gatewayEventStoreRepository struct {
		MongoDB mongodbInfrastructure.IMongoDBInfrastructure
		logger  pkg.IZapLogger
	}
)

func NewGatewayEventStoreRepository(mongodb mongodbInfrastructure.IMongoDBInfrastructure, logger pkg.IZapLogger) IGatewayEventStoreRepository {
	return &gatewayEventStoreRepository{
		MongoDB: mongodb,
		logger:  logger,
	}
}
