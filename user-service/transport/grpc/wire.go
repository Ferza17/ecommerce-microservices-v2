//go:build wireinject
// +build wireinject

package grpc

import (
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/postgres"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/redis"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
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

		// Repository Layer
		userPostgresqlRepository.Set,
		userRedisRepository.Set,
		rolePostgresqlRepository.Set,

		// UseCase Layer
		userUseCase.Set,
		authUseCase.Set,

		// Presenter Layer
		authPresenter.Set,
		userPresenter.Set,

		Set,
	)
	return nil
}
