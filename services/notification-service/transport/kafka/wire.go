//go:build wireinject
// +build wireinject

package kafka

import (
	kafkaInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/kafka"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/mailhog"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/mongodb"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/rabbitmq"
	paymentService "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/services/payment"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	"github.com/google/wire"

	notificationEmailConsumer "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/email/consumer"
	notificationEmailMongoDBRepository "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/email/repository/mongodb"
	notificationEmailUseCase "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/email/usecase"
)

func ProvideServer() *Server {
	wire.Build(
		logger.Set,
		kafkaInfrastructure.Set,
		mailhog.Set,
		mongodb.Set,
		telemetry.Set,
		rabbitmq.Set,
		paymentService.Set,

		notificationEmailMongoDBRepository.Set,
		notificationEmailUseCase.Set,
		notificationEmailConsumer.Set,

		Set,
	)
	return nil
}
