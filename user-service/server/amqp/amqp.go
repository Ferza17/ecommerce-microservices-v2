package amqp

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure"
	userConsumer "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/consumer"
	userPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/repository/postgresql"
	userUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/usecase"
	eventStoreRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/userEventStore/repository/mongodb"
	eventStoreUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/userEventStore/usecase"
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
		amqpConn *amqp091.Connection

		logger              pkg.IZapLogger
		postgresqlConnector *infrastructure.PostgresqlConnector
		mongoDBConnector    *infrastructure.MongodbConnector
	}

	Option func(server *Server)
)

func NewServer(option ...Option) *Server {

	amqpConn, err := amqp091.Dial(
		fmt.Sprintf("amqp://%s:%s@%s:%s/",
			config.Get().RabbitMQUsername,
			config.Get().RabbitMQPassword,
			config.Get().RabbitMQHost,
			config.Get().RabbitMQPort,
		))
	if err != nil {
		log.Fatalf("error while connecting to RabbitMQ: %v\n", err)
	}
	log.Println("RabbitMQ connected")

	s := &Server{
		amqpConn: amqpConn,
	}
	for _, o := range option {
		o(s)
	}
	return s
}

func (srv *Server) Serve() {
	amqpChannel, err := srv.amqpConn.Channel()
	if err != nil {
		srv.logger.Error(fmt.Sprintf("failed to serve", zap.Error(err)))
	}

	if err := amqpChannel.ExchangeDeclare(
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

	// Register Module

	newUserEventStoreRepository := eventStoreRepository.NewEventStoreRepository(srv.mongoDBConnector, srv.logger)
	newUserEventStoreUseCase := eventStoreUseCase.NewUserEventStoreUseCase(newUserEventStoreRepository, srv.logger)

	newUserPostgresqlRepository := userPostgresqlRepository.NewUserPostgresqlRepository(srv.postgresqlConnector, srv.logger)
	newUserUseCase := userUseCase.NewUserUseCase(newUserPostgresqlRepository, newUserEventStoreUseCase, srv.logger)
	newUserConsumer := userConsumer.NewUserConsumer(amqpChannel, newUserUseCase, srv.logger)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		log.Println("AMQP shutdown...")
	}()

	go func() {
		if err = newUserConsumer.UserCreated(ctx); err != nil {
			srv.logger.Error(fmt.Sprintf("failed to UserCreated", zap.Error(err)))
		}
	}()

	go func() {
		if err = newUserConsumer.UserUpdated(ctx); err != nil {
			srv.logger.Error(fmt.Sprintf("failed to UserUpdated", zap.Error(err)))
		}
	}()

	<-ctx.Done()
}
