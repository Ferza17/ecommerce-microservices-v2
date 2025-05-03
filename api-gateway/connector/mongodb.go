package connector

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type (
	IMongodbConnector interface {
		Close(ctx context.Context) error
	}

	MongodbConnector struct {
		MongoClient *mongo.Client
	}
)

func NewMongodbConnector() IMongodbConnector {
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
		log.Fatalf("error while connecting to MongoDB: %v\n", err)
	}
	// Make sure that connection insurable
	if err = conn.Ping(context.Background(), nil); err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	}
	log.Println("MongoDB connected")
	return &MongodbConnector{
		MongoClient: conn,
	}
}

func (m *MongodbConnector) Close(ctx context.Context) error {
	return m.MongoClient.Disconnect(ctx)
}
