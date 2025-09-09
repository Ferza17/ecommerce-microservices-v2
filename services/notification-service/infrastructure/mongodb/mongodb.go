package mongodb

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/enum"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type (
	IMongoDBInfrastructure interface {
		GetCollection(database enum.Database, collection enum.Collection) *mongo.Collection
		GetConnectionString() string
		Close(ctx context.Context) error
		StartSession() (mongo.Session, error)
	}

	MongoDBInfrastructure struct {
		mongoClient *mongo.Client
		logger      logger.IZapLogger
	}
)

var Set = wire.NewSet(NewMongoDBInfrastructure)

func NewMongoDBInfrastructure(logger logger.IZapLogger) IMongoDBInfrastructure {
	conn, err := mongo.Connect(
		context.Background(),
		options.
			Client().
			ApplyURI(
				fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin",
					config.Get().MongoUsername,
					config.Get().MongoPassword,
					config.Get().MongoHost,
					config.Get().MongoPort,
					config.Get().MongoDatabaseName,
				),
			),
	)

	if err != nil {
		logger.Error("MongoDBInfrastructure.NewMongoDBInfrastructure", zap.Error(err))
	}
	// Make sure that connection insurable
	if err = conn.Ping(context.Background(), nil); err != nil {
		logger.Error("MongoDBInfrastructure.NewMongoDBInfrastructure", zap.Error(err))
	}
	return &MongoDBInfrastructure{
		mongoClient: conn,
		logger:      logger,
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

func (m *MongoDBInfrastructure) GetCollection(database enum.Database, collection enum.Collection) *mongo.Collection {
	return m.mongoClient.Database(database.String()).Collection(collection.String())
}

func (m *MongoDBInfrastructure) GetConnectionString() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin",
		config.Get().MongoUsername,
		config.Get().MongoPassword,
		config.Get().MongoHost,
		config.Get().MongoPort,
		config.Get().MongoDatabaseName,
	)
}
