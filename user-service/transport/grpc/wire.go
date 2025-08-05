//go:build wireinject
// +build wireinject

package grpc

import (
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/postgres"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/redis"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/temporal"
	accessControlPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/accessControl/repository/postgres"
	accessControlRedisRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/accessControl/repository/redis"
	accessControlUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/accessControl/usecase"
	authPresenter "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/presenter"
	userRedisRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/repository/redis"
	authUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/usecase"
	rolePostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/role/repository/postgres"
	userPresenter "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/presenter"
	userPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/repository/postgres"
	userUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/usecase"

	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
)

func ProvideGrpcServer() *Server {
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
		authPresenter.Set,
		userPresenter.Set,

		Set,
	)
	return nil
}
