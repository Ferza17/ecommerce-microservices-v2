package kafka

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	kafkaInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/kafka"
	roleUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/role/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IRoleConsumer interface {
		DlqSinkPgUsersRoles(ctx context.Context, message *kafka.Message) error
	}

	roleConsumer struct {
		kafkaInfrastructure kafkaInfrastructure.IKafkaInfrastructure
		logger              logger.IZapLogger
		roleUseCase         roleUseCase.IRoleUseCase
	}
)

var Set = wire.NewSet(NewRoleConsumer)

func NewRoleConsumer(
	kafkaInfrastructure kafkaInfrastructure.IKafkaInfrastructure,
	logger logger.IZapLogger,
	roleUseCase roleUseCase.IRoleUseCase,
) IRoleConsumer {
	return &roleConsumer{
		kafkaInfrastructure: kafkaInfrastructure,
		logger:              logger,
		roleUseCase:         roleUseCase,
	}
}
