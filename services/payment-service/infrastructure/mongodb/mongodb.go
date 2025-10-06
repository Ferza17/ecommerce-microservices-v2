package mongodb

import (
	"context"
	"fmt"

	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type (
	IMongoDBInfrastructure interface {
		GetCollection(collection string) *mongo.Collection
		GetConnectionString() string
		Close(ctx context.Context) error
		StartSession() (mongo.Session, error)
	}

	MongoDBInfrastructure struct {
		mongoClient *mongo.Client
		logger      logger.IZapLogger
		database    string
	}
)

var Set = wire.NewSet(NewMongoDBInfrastructure)

func NewMongoDBInfrastructure(logger logger.IZapLogger) IMongoDBInfrastructure {
	conUrl := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/%s?directConnection=true&replicaSet=rs0",
		config.Get().DatabaseMongo.Username,
		config.Get().DatabaseMongo.Password,
		config.Get().DatabaseMongo.Host,
		config.Get().DatabaseMongo.Port,
		config.Get().DatabaseMongo.DatabaseName,
	)
	conn, err := mongo.Connect(
		context.Background(),
		options.
			Client().
			ApplyURI(
				conUrl,
			),
	)
	if err != nil {
		logger.Error("MongoDBInfrastructure.NewMongoDBInfrastructure", zap.Error(err))
		return nil
	}

	if err = conn.Ping(context.Background(), nil); err != nil {
		logger.Error("MongoDBInfrastructure.NewMongoDBInfrastructure", zap.Error(err))
		return nil
	}

	return &MongoDBInfrastructure{
		mongoClient: conn,
		logger:      logger,
		database:    config.Get().DatabaseMongo.DatabaseName,
	}
}

func (m *MongoDBInfrastructure) Close(ctx context.Context) error {
	if err := m.mongoClient.Disconnect(ctx); err != nil {
		m.logger.Error("MongoDBInfrastructure.Close", zap.Error(err))
		return err
	}
	return nil
}

func (m *MongoDBInfrastructure) StartSession() (mongo.Session, error) {
	return m.mongoClient.StartSession()
}

func (m *MongoDBInfrastructure) GetCollection(collection string) *mongo.Collection {
	return m.mongoClient.Database(m.database).Collection(collection)
}

func (m *MongoDBInfrastructure) GetConnectionString() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?directConnection=true&replicaSet=rs0",
		config.Get().DatabaseMongo.Username,
		config.Get().DatabaseMongo.Password,
		config.Get().DatabaseMongo.Host,
		config.Get().DatabaseMongo.Port,
		config.Get().DatabaseMongo.DatabaseName,
	)
}
