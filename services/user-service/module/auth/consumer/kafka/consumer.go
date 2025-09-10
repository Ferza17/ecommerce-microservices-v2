package kafka

import (
	"context"
	kafkaInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/kafka"
	authUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IAuthConsumer interface {
		UserLogin(ctx context.Context) error
	}

	authConsumer struct {
		kafkaInfrastructure kafkaInfrastructure.KafkaInfrastructure
		logger              logger.IZapLogger
		authUseCase         authUseCase.IAuthUseCase
	}
)

var Set = wire.NewSet(NewAuthConsumer)

func NewAuthConsumer(
	kafkaInfrastructure kafkaInfrastructure.KafkaInfrastructure,
	logger logger.IZapLogger,
	authUseCase authUseCase.IAuthUseCase) IAuthConsumer {
	return &authConsumer{
		kafkaInfrastructure: kafkaInfrastructure,
		logger:              logger,
		authUseCase:         authUseCase,
	}
}
