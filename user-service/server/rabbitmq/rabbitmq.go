package rabbitmq

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/bootstrap"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
	authConsumer "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/consumer"
	authUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/usecase"
	userConsumer "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/consumer"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type (
	Server struct {
		amqpInfrastructure rabbitmq.IRabbitMQInfrastructure
		logger             pkg.IZapLogger
		userUseCase        usecase.IUserUseCase
		authUseCase        authUseCase.IAuthUseCase
		userConsumer       userConsumer.IUserConsumer
		authConsumer       authConsumer.IAuthConsumer
	}
)

func NewServer(dependency *bootstrap.Bootstrap) *Server {
	return &Server{
		amqpInfrastructure: dependency.RabbitMQInfrastructure,
		userUseCase:        dependency.UserUseCase,
		authUseCase:        dependency.AuthUseCase,
		logger:             dependency.Logger,
		userConsumer:       dependency.UserConsumer,
		authConsumer:       dependency.AuthConsumer,
	}
}

func (srv *Server) Serve() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		log.Println("AMQP shutdown...")
	}()

	go func() {
		if err := srv.userConsumer.UserCreated(ctx); err != nil {
			srv.logger.Error(fmt.Sprintf("failed to UserCreated", zap.Error(err)))
		}
	}()

	go func() {
		if err := srv.userConsumer.UserUpdated(ctx); err != nil {
			srv.logger.Error(fmt.Sprintf("failed to UserUpdated", zap.Error(err)))
		}
	}()

	go func() {
		if err := srv.authConsumer.UserLogin(ctx); err != nil {
			srv.logger.Error(fmt.Sprintf("failed to UserLogin", zap.Error(err)))
		}
	}()

	<-ctx.Done()
}
