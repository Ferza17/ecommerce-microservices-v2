package bootstrap

import (
	mongoDBInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/mongodb"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg"
)

type Bootstrap struct {
	Logger pkg.IZapLogger

	MongoDBInfrastructure  mongoDBInfrastructure.IMongoDBInfrastructure
	RabbitMQInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure
}

func NewBootstrap() *Bootstrap {
	logger := pkg.NewZapLogger()

	newMongoDBInfrastructure := mongoDBInfrastructure.NewMongoDBInfrastructure(logger)
	newRabbitMQInfrastructure := rabbitmqInfrastructure.NewRabbitMQInfrastructure(logger)

	return &Bootstrap{
		Logger:                 logger,
		MongoDBInfrastructure:  newMongoDBInfrastructure,
		RabbitMQInfrastructure: newRabbitMQInfrastructure,
	}
}
