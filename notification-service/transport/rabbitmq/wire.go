//go:build wireinject
// +build wireinject

package rabbitmq

import (
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/mailhog"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/mongodb"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	"github.com/google/wire"

	notificationEmailConsumer "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/email/consumer"
	notificationEmailMongoDBRepository "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/email/repository/mongodb"
	notificationEmailUseCase "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/email/usecase"
)

func ProvideRabbitMQServer() *RabbitMQTransport {
	wire.Build(
		logger.Set,
		mailhog.Set,
		mongodb.Set,
		telemetry.Set,
		rabbitmq.Set,

		notificationEmailMongoDBRepository.Set,
		notificationEmailUseCase.Set,
		notificationEmailConsumer.Set,

		Set,
	)
	return nil
}
