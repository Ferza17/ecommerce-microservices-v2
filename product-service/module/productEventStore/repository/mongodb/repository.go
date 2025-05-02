package mongodb

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/connector"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/bson"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	IProductEventStoreRepository interface {
		CreateProductEventStore(ctx context.Context, requestId string, req *bson.Event) (string, error)
		CreateProductEventStoreWithSession(ctx context.Context, requestId string, req *bson.Event, session mongo.Session) (string, error)

		StartSession() (mongo.Session, error)
	}

	productEventStoreRepository struct {
		connector *connector.MongodbConnector
		logger    pkg.IZapLogger
	}
)

const (
	databaseEventStore     = "event-store"
	collectionProductEvent = "event-store-product-event"
)

func NewProductEventStoreRepository(connector *connector.MongodbConnector, logger pkg.IZapLogger) IProductEventStoreRepository {
	return &productEventStoreRepository{
		connector: connector,
		logger:    logger,
	}
}

func (r *productEventStoreRepository) StartSession() (mongo.Session, error) {
	return r.connector.MongoClient.StartSession()
}
