package rabbitmq

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
	authConsumer "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/consumer"
	userConsumer "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/consumer"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
	"go.uber.org/zap"
	"sync"
)

type (
	Server struct {
		amqpInfrastructure rabbitmq.IRabbitMQInfrastructure
		logger             logger.IZapLogger
		userConsumer       userConsumer.IUserConsumer
		authConsumer       authConsumer.IAuthConsumer
	}
)

var Set = wire.NewSet(NewServer)

func NewServer(
	amqpInfrastructure rabbitmq.IRabbitMQInfrastructure,
	logger logger.IZapLogger,
	userConsumer userConsumer.IUserConsumer,
	authConsumer authConsumer.IAuthConsumer,
) *Server {
	return &Server{
		amqpInfrastructure: amqpInfrastructure,
		logger:             logger,
		userConsumer:       userConsumer,
		authConsumer:       authConsumer,
	}
}

func (srv *Server) Serve(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := srv.userConsumer.UserCreated(ctx); err != nil {
			srv.logger.Error("User Created Error", zap.Error(err))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := srv.userConsumer.UserUpdated(ctx); err != nil {
			srv.logger.Error("User Updated Error", zap.Error(err))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := srv.authConsumer.UserLogin(ctx); err != nil {
			srv.logger.Error("User Login Error", zap.Error(err))
		}
	}()

	wg.Wait()
	cancel()
}
