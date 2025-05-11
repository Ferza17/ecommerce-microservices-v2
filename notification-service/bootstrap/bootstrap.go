package bootstrap

import (
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/mailhog"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/mongodb"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/rabbitmq"
	notificationConsumer "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/notification/consumer"
	notificationRepository "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/notification/repository/mongodb"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/module/notification/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg"
)

type Bootstrap struct {
	Logger                 pkg.IZapLogger
	RabbitMQInfrastructure rabbitmq.IRabbitMQInfrastructure
	MongoDBInfrastructure  mongodb.IMongoDBInfrastructure
	MailHogInfrastructure  mailhog.IMailhogInfrastructure
	NotificationRepository notificationRepository.INotificationRepository
	NotificationUseCase    usecase.INotificationUseCase
	NotificationConsumer   notificationConsumer.INotificationConsumer
}

func NewBootstrap() *Bootstrap {
	logger := pkg.NewZapLogger()

	// Infrastructure
	rabbitmqInfrastructure := rabbitmq.NewRabbitMQInfrastructure(logger)
	mongodbInfrastructure := mongodb.NewMongoDBInfrastructure(logger)
	mailHogInfrastructure := mailhog.NewMailhogInfrastructure(logger)

	// Repository
	notificationRepository := notificationRepository.NewEventRepository(mongodbInfrastructure, logger)

	// UseCase
	notificationUseCase := usecase.NewEventStoreUseCase(notificationRepository, rabbitmqInfrastructure, mailHogInfrastructure, logger)

	// Consumer
	notificationConsumer := notificationConsumer.NewNotificationConsumer(rabbitmqInfrastructure, notificationUseCase, logger)

	return &Bootstrap{
		Logger:                 logger,
		RabbitMQInfrastructure: rabbitmqInfrastructure,
		MongoDBInfrastructure:  mongodbInfrastructure,
		MailHogInfrastructure:  mailHogInfrastructure,
		NotificationRepository: notificationRepository,
		NotificationUseCase:    notificationUseCase,
		NotificationConsumer:   notificationConsumer,
	}
}
