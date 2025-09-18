//go:build wireinject
// +build wireinject

package kafka

import (
	kafkaInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/kafka"
	postgresInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/postgres"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
	redisInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/redis"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	accessControlPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/accessControl/repository/postgres"
	accessControlRedisRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/accessControl/repository/redis"
	accessControlUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/accessControl/usecase"
	authKafkaConsumer "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/consumer/kafka"
	authRedisRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/repository/redis"
	authUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/usecase"
	rolePostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/role/repository/postgres"
	userKafkaConsumer "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/consumer/kafka"
	userPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/repository/postgres"
	userUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
)

func ProvideServer() *Server {
	wire.Build(
		logger.Set,

		kafkaInfrastructure.Set,
		telemetryInfrastructure.Set,
		rabbitmqInfrastructure.Set,
		postgresInfrastructure.Set,
		redisInfrastructure.Set,

		authKafkaConsumer.Set,
		userKafkaConsumer.Set,

		userPostgresqlRepository.Set,
		rolePostgresqlRepository.Set,
		accessControlPostgresqlRepository.Set,
		accessControlRedisRepository.Set,
		authRedisRepository.Set,

		authUseCase.Set,
		userUseCase.Set,
		accessControlUseCase.Set,

		Set,
	)
	return nil
}
