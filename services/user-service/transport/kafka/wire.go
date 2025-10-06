//go:build wireinject
// +build wireinject

package kafka

import (
	kafkaInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/kafka"
	mongoDBInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/mongodb"
	postgresInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/postgres"
	redisInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/redis"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	accessControlPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/accessControl/repository/postgres"
	accessControlRedisRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/accessControl/repository/redis"
	accessControlUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/accessControl/usecase"
	authKafkaConsumer "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/consumer/kafka"
	authRedisRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/repository/redis"
	authUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/usecase"
	eventKafkaConsumer "github.com/ferza17/ecommerce-microservices-v2/user-service/module/event/consumer"
	eventMongoRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/event/repository/mongodb"
	eventUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/event/usecase"
	rolesKafkaConsumer "github.com/ferza17/ecommerce-microservices-v2/user-service/module/role/consumer/kafka"
	rolePostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/role/repository/postgres"
	roleUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/role/usecase"
	userKafkaConsumer "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/consumer/kafka"
	userPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/repository/postgres"
	userUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/usecase"

	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
)

func Provide() *Transport {
	wire.Build(
		logger.Set,

		kafkaInfrastructure.Set,
		telemetryInfrastructure.Set,
		postgresInfrastructure.Set,
		redisInfrastructure.Set,
		mongoDBInfrastructure.Set,

		authKafkaConsumer.Set,
		userKafkaConsumer.Set,
		rolesKafkaConsumer.Set,
		eventKafkaConsumer.Set,

		userPostgresqlRepository.Set,
		rolePostgresqlRepository.Set,
		accessControlPostgresqlRepository.Set,
		accessControlRedisRepository.Set,
		authRedisRepository.Set,
		eventMongoRepository.Set,

		authUseCase.Set,
		userUseCase.Set,
		accessControlUseCase.Set,
		roleUseCase.Set,
		eventUseCase.Set,

		Set,
	)
	return nil
}
