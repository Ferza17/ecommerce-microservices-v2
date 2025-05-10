package bootstrap

import (
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/postgresql"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/redis"
	authConsumer "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/consumer"
	authRedisRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/repository/redis"
	authUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/usecase"
	userConsumer "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/consumer"
	userPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/repository/postgresql"
	userUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
)

type Bootstrap struct {
	Logger                   pkg.IZapLogger
	RabbitMQInfrastructure   rabbitmq.IRabbitMQInfrastructure
	PostgresqlInfrastructure postgresql.IPostgreSQLInfrastructure
	RedisInfrastructure      redis.IRedisInfrastructure
	UserPostgresqlRepository userPostgresqlRepository.IUserPostgresqlRepository
	UserUseCase              userUseCase.IUserUseCase
	AuthUseCase              authUseCase.IAuthUseCase
	AuthRedisRepository      authRedisRepository.IAuthRedisRepository
	AuthConsumer             authConsumer.IAuthConsumer
	UserConsumer             userConsumer.IUserConsumer
}

func NewBootstrap() *Bootstrap {
	logger := pkg.NewZapLogger()

	// Infrastructure
	newRabbitMQInfrastructure := rabbitmq.NewRabbitMQInfrastructure(logger)
	newPostgresqlInfrastructure := postgresql.NewPostgresqlInfrastructure(logger)
	newRedisInfrastructure := redis.NewRedisInfrastructure(logger)

	// Repository
	newUserPostgresqlRepository := userPostgresqlRepository.NewUserPostgresqlRepository(newPostgresqlInfrastructure, logger)
	newAuthRedisRepository := authRedisRepository.NewAuthRedisRepository(newRedisInfrastructure, logger)

	// usecase
	newUserUseCase := userUseCase.NewUserUseCase(newUserPostgresqlRepository, newRabbitMQInfrastructure, logger)
	newAuthUseCase := authUseCase.NewAuthUseCase(newUserPostgresqlRepository, newAuthRedisRepository, newRabbitMQInfrastructure, logger)

	// Consumer
	newUserConsumer := userConsumer.NewUserConsumer(newRabbitMQInfrastructure, newUserUseCase, logger)
	newAuthConsumer := authConsumer.NewAuthConsumer(newRabbitMQInfrastructure, newAuthUseCase, logger)

	return &Bootstrap{
		Logger:                   logger,
		UserUseCase:              newUserUseCase,
		AuthUseCase:              newAuthUseCase,
		UserPostgresqlRepository: newUserPostgresqlRepository,
		PostgresqlInfrastructure: newPostgresqlInfrastructure,
		RedisInfrastructure:      newRedisInfrastructure,
		RabbitMQInfrastructure:   newRabbitMQInfrastructure,
		AuthRedisRepository:      newAuthRedisRepository,
		AuthConsumer:             newAuthConsumer,
		UserConsumer:             newUserConsumer,
	}
}
