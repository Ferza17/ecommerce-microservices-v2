package bootstrap

import (
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/postgresql"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
	userPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/repository/postgresql"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
)

type Bootstrap struct {
	Logger                   pkg.IZapLogger
	RabbitMQInfrastructure   rabbitmq.IRabbitMQInfrastructure
	PostgresqlInfrastructure postgresql.IPostgreSQLInfrastructure
	UserPostgresqlRepository userPostgresqlRepository.IUserPostgresqlRepository
	UserUseCase              usecase.IUserUseCase
}

func NewBootstrap() *Bootstrap {
	logger := pkg.NewZapLogger()

	// Infrastructure
	newRabbitMQInfrastructure := rabbitmq.NewRabbitMQInfrastructure(logger)
	newPostgresqlInfrastructure := postgresql.NewPostgresqlInfrastructure(logger)

	// Repository
	newUserPostgresqlRepository := userPostgresqlRepository.NewUserPostgresqlRepository(newPostgresqlInfrastructure, logger)

	// usecase
	newUserUseCase := usecase.NewUserUseCase(newUserPostgresqlRepository, newRabbitMQInfrastructure, logger)

	return &Bootstrap{
		Logger:                   logger,
		UserUseCase:              newUserUseCase,
		UserPostgresqlRepository: newUserPostgresqlRepository,
		PostgresqlInfrastructure: newPostgresqlInfrastructure,
		RabbitMQInfrastructure:   newRabbitMQInfrastructure,
	}
}
