package rabbitmq

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/bootstrap"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
	userConsumer "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/consumer"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type (
	Server struct {
		amqpConn    *amqp091.Connection
		logger      pkg.IZapLogger
		userUseCase usecase.IUserUseCase
	}
)

func NewServer(dependency *bootstrap.Bootstrap) *Server {
	amqpConn, err := amqp091.Dial(
		fmt.Sprintf("amqp://%s:%s@%s:%s/",
			config.Get().RabbitMQUsername,
			config.Get().RabbitMQPassword,
			config.Get().RabbitMQHost,
			config.Get().RabbitMQPort,
		))
	if err != nil {
		dependency.Logger.Error(fmt.Sprintf("Failed to connect to RabbitMQ: %v", err))
	}
	log.Println("RabbitMQ connected")
	return &Server{
		amqpConn:    amqpConn,
		userUseCase: dependency.UserUseCase,
		logger:      dependency.Logger,
	}
}

func (srv *Server) Serve() {
	amqpChannel, err := srv.amqpConn.Channel()
	if err != nil {
		srv.logger.Error(fmt.Sprintf("failed to serve", zap.Error(err)))
	}

	if err = amqpChannel.ExchangeDeclare(
		enum.USER_EXCHANGE.String(),
		"topic", // type
		true,    // durable
		false,   // auto-delete
		false,
		true,
		nil,
	); err != nil {
		srv.logger.Error(fmt.Sprintf("failed to serve", zap.Error(err)))
	}

	userConsumer := userConsumer.NewUserConsumer(amqpChannel, srv.userUseCase, srv.logger)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		log.Println("AMQP shutdown...")
	}()

	go func() {
		if err = userConsumer.UserCreated(ctx); err != nil {
			srv.logger.Error(fmt.Sprintf("failed to UserCreated", zap.Error(err)))
		}
	}()

	go func() {
		if err = userConsumer.UserUpdated(ctx); err != nil {
			srv.logger.Error(fmt.Sprintf("failed to UserUpdated", zap.Error(err)))
		}
	}()

	<-ctx.Done()

	if err = amqpChannel.Close(); err != nil {
		srv.logger.Error(fmt.Sprintf("failed to close channel %v", zap.Error(err)))
		return
	}
}
