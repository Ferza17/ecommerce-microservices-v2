package mongodb

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/enum"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	IMongoDBInfrastructure interface {
		GetCollection(database enum.Database, collection enum.Collection) *mongo.Collection
		Close(ctx context.Context) error
		StartSession() (mongo.Session, error)
	}

	MongoDBInfrastructure struct {
		mongoClient *mongo.Client
		logger      pkg.IZapLogger
	}
)

func NewMongoDBInfrastructure(logger pkg.IZapLogger) IMongoDBInfrastructure {
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
		logger.Error(fmt.Sprintf("Could not connect to MongoDB: %v\n", err))
	}
	// Make sure that connection insurable
	if err = conn.Ping(context.Background(), nil); err != nil {
		logger.Error(fmt.Sprintf("Could not ping MongoDB: %v\n", err))
	}
	return &MongoDBInfrastructure{
		mongoClient: conn,
		logger:      logger,
	}
}

func (m *MongoDBInfrastructure) Close(ctx context.Context) error {
	if err := m.mongoClient.Disconnect(ctx); err != nil {
		m.logger.Error(fmt.Sprintf("Failed to close a connection: %v", err))
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
