//go:build wireinject
// +build wireinject

package rabbitmq

import (
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/postgres"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/redis"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/temporal"
	accessControlPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/accessControl/repository/postgres"
	accessControlRedisRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/accessControl/repository/redis"
	accessControlUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/accessControl/usecase"
	authConsumer "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/consumer"
	userRedisRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/repository/redis"
	authUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/usecase"
	authWorkflow "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/workflow"
	rolePostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/role/repository/postgres"
	userConsumer "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/consumer"
	userPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/repository/postgres"

	userUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
)

func ProvideRabbitMQServer() *Server {
	wire.Build(
		logger.Set,

		// Infrastructure Layer
		redis.Set,
		postgres.Set,
		rabbitmq.Set,
		telemetry.Set,
		temporal.Set,

		// Repository Layer
		userPostgresqlRepository.Set,
		userRedisRepository.Set,
		rolePostgresqlRepository.Set,
		accessControlPostgresqlRepository.Set,
		accessControlRedisRepository.Set,

		// UseCase Layer
		userUseCase.Set,
		authUseCase.Set,
		accessControlUseCase.Set,

		// Presenter Layer
		authConsumer.Set,
		userConsumer.Set,

		// Workflow Layer
		authWorkflow.Set,

		Set,
	)
	return nil
}
