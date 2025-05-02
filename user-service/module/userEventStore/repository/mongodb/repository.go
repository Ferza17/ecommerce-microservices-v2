package mongodb

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/connector"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/bson"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	IUserEventStoreRepository interface {
		CreateUserEventStore(ctx context.Context, requestId string, req *bson.Event) (string, error)
		CreateUserEventStoreWithSession(ctx context.Context, requestId string, event *bson.Event, session mongo.Session) (string, error)

		// Session
		StartSession(ctx context.Context) (mongo.Session, error)
	}

	userEventStoreRepository struct {
		connector *connector.MongodbConnector
		logger    pkg.IZapLogger
	}
)

func (r *userEventStoreRepository) StartSession(ctx context.Context) (mongo.Session, error) {
	return r.connector.MongoClient.StartSession()
}

const (
	databaseEventStore  = "event-store"
	collectionUserEvent = "event-store-user-event"
)

func NewEventStoreRepository(connector *connector.MongodbConnector, logger pkg.IZapLogger) IUserEventStoreRepository {
	return &userEventStoreRepository{
		connector: connector,
		logger:    logger,
	}
}
