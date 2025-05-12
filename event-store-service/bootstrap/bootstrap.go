package bootstrap

import (
	mongoDBInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/mongodb"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/module/event/consumer"
	eventRepository "github.com/ferza17/ecommerce-microservices-v2/event-store-service/module/event/repository/mongodb"
	eventUseCase "github.com/ferza17/ecommerce-microservices-v2/event-store-service/module/event/usecase"

	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg"
)

type Bootstrap struct {
	Logger pkg.IZapLogger

	TelemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
	MongoDBInfrastructure   mongoDBInfrastructure.IMongoDBInfrastructure
	RabbitMQInfrastructure  rabbitmqInfrastructure.IRabbitMQInfrastructure

	EventRepository eventRepository.IEventRepository
	EventUseCase    eventUseCase.IEventUseCase
	EventConsumer   consumer.IEventConsumer
}

func NewBootstrap() *Bootstrap {
	logger := pkg.NewZapLogger()

	newTelemetryInfrastructure := telemetryInfrastructure.NewTelemetry(logger)
	newMongoDBInfrastructure := mongoDBInfrastructure.NewMongoDBInfrastructure(logger)
	newRabbitMQInfrastructure := rabbitmqInfrastructure.NewRabbitMQInfrastructure(logger)

	// Register Repository , UseCase, Consumer
	newEventRepository := eventRepository.NewEventRepository(newMongoDBInfrastructure, newTelemetryInfrastructure, logger)
	newEventUseCase := eventUseCase.NewEventStoreUseCase(newEventRepository, newRabbitMQInfrastructure, newTelemetryInfrastructure, logger)
	newEventConsumer := consumer.NewEventConsumer(newRabbitMQInfrastructure, newTelemetryInfrastructure, newEventUseCase, logger)

	return &Bootstrap{
		Logger:                  logger,
		TelemetryInfrastructure: newTelemetryInfrastructure,
		MongoDBInfrastructure:   newMongoDBInfrastructure,
		RabbitMQInfrastructure:  newRabbitMQInfrastructure,
		EventRepository:         newEventRepository,
		EventUseCase:            newEventUseCase,
		EventConsumer:           newEventConsumer,
	}
}
