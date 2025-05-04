package mongodb

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/bson"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	IUserEventStoreRepository interface {
		CreateUserEventStore(ctx context.Context, requestId string, req *bson.Event) (string, error)
		CreateUserEventStoreWithSession(ctx context.Context, requestId string, event *bson.Event, session mongo.Session) (string, error)

		// Session
		StartSession() (mongo.Session, error)
	}

	userEventStoreRepository struct {
		connector *infrastructure.MongodbConnector
		logger    pkg.IZapLogger
	}
)

const (
	databaseEventStore  = "event-store"
	collectionUserEvent = "event-store-user-event"
)

func NewEventStoreRepository(connector *infrastructure.MongodbConnector, logger pkg.IZapLogger) IUserEventStoreRepository {
	return &userEventStoreRepository{
		connector: connector,
		logger:    logger,
	}
}

func (r *userEventStoreRepository) StartSession() (mongo.Session, error) {
	return r.connector.MongoClient.StartSession()
}
