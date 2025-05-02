package connector

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongodbConnector struct {
	mongoClient *mongo.Client
}

func NewMongodbConnector() *MongodbConnector {
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
		mongoClient: conn,
	}
}

func (m *MongodbConnector) Close(ctx context.Context) error {
	return m.mongoClient.Disconnect(ctx)
}
