package consumer

import (
	"context"
	userUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
	"github.com/rabbitmq/amqp091-go"
)

type (
	IUserConsumer interface {
		UserCreated(ctx context.Context) error
		UserUpdated(ctx context.Context) error
	}
	userConsumer struct {
		amqpChannel *amqp091.Channel
		userUseCase userUseCase.IUserUseCase
		logger      pkg.IZapLogger
	}
)

func NewUserConsumer(amqpChannel *amqp091.Channel, userUseCase userUseCase.IUserUseCase, logger pkg.IZapLogger) IUserConsumer {
	return &userConsumer{
		amqpChannel: amqpChannel,
		userUseCase: userUseCase,
		logger:      logger,
	}
}
