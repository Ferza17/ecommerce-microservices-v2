//go:build wireinject
// +build wireinject

package rabbitmq

import (
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/mailhog"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/mongodb"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/rabbitmq"
	paymentService "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/services/payment"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/temporal"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	"github.com/google/wire"

	notificationEmailConsumer "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/email/consumer"
	notificationEmailMongoDBRepository "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/email/repository/mongodb"
	notificationEmailUseCase "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/email/usecase"
	notificationEmailWorkflow "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/email/workflow"
)

func ProvideRabbitMQServer() *RabbitMQTransport {
	wire.Build(
		logger.Set,
		mailhog.Set,
		mongodb.Set,
		telemetry.Set,
		rabbitmq.Set,
		paymentService.Set,
		temporal.Set,

		notificationEmailMongoDBRepository.Set,
		notificationEmailUseCase.Set,
		notificationEmailConsumer.Set,
		notificationEmailWorkflow.Set,

		Set,
	)
	return nil
}
