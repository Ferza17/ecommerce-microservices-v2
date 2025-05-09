package rabbitmq

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/bootstrap"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
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
	}
)

func NewServer(dependency *bootstrap.Bootstrap) *Server {
	return &Server{
		amqpInfrastructure: dependency.RabbitMQInfrastructure,
		userUseCase:        dependency.UserUseCase,
		logger:             dependency.Logger,
	}
}

func (srv *Server) Serve() {
	userConsumer := userConsumer.NewUserConsumer(srv.amqpInfrastructure, srv.userUseCase, srv.logger)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		log.Println("AMQP shutdown...")
	}()

	go func() {
		if err := userConsumer.UserCreated(ctx); err != nil {
			srv.logger.Error(fmt.Sprintf("failed to UserCreated", zap.Error(err)))
		}
	}()

	go func() {
		if err := userConsumer.UserUpdated(ctx); err != nil {
			srv.logger.Error(fmt.Sprintf("failed to UserUpdated", zap.Error(err)))
		}
	}()

	<-ctx.Done()
}
