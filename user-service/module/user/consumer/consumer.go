package consumer

import (
	userUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
	"github.com/rabbitmq/amqp091-go"
)

type UserConsumer struct {
	amqpChannel *amqp091.Channel
	UserUseCase userUseCase.IUserUseCase
	Logger      pkg.IZapLogger
}

func NewUserConsumer(userUseCase userUseCase.IUserUseCase, amqpChannel *amqp091.Channel, logger pkg.IZapLogger) *UserConsumer {
	return &UserConsumer{
		UserUseCase: userUseCase,
		Logger:      logger,
		amqpChannel: amqpChannel,
	}
}
