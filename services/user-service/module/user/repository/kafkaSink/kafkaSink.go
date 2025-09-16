package kafkaSink

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	kafkaInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/kafka"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IUserKafkaSink interface {
		CreateUser(ctx context.Context, requestId string, req *orm.User) error
	}

	userKafkaSink struct {
		logger              logger.IZapLogger
		kafkaInfrastructure kafkaInfrastructure.IKafkaInfrastructure
	}
)

func (r *userKafkaSink) CreateUser(ctx context.Context, requestId string, req *orm.User) error {
	if err := r.kafkaInfrastructure.PublishWithJsonSchema(ctx, config.Get().BrokerKafkaTopicConnectorSinkPgUser.Users, req.ID, req); err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error creating user: %v", requestId, err))
		return err
	}
	return nil
}

var Set = wire.NewSet(NewUserKafkaSink)

func NewUserKafkaSink(
	logger logger.IZapLogger,
	kafkaInfrastructure kafkaInfrastructure.IKafkaInfrastructure,
) IUserKafkaSink {
	return &userKafkaSink{
		logger:              logger,
		kafkaInfrastructure: kafkaInfrastructure,
	}
}
