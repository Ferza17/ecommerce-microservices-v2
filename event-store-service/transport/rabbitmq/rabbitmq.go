package rabbitmq

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/config"
	mongoDBInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/mongodb"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/module/event/consumer"
	eventRepository "github.com/ferza17/ecommerce-microservices-v2/event-store-service/module/event/repository/mongodb"
	eventUseCase "github.com/ferza17/ecommerce-microservices-v2/event-store-service/module/event/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type (
	RabbitMQTransport struct {
		amqpConn *amqp091.Connection

		logger                pkg.IZapLogger
		mongoDBInfrastructure mongoDBInfrastructure.IMongoDBInfrastructure
	}

	Option func(server *RabbitMQTransport)
)

func NewServer(option ...Option) *RabbitMQTransport {
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

	s := &RabbitMQTransport{
		amqpConn: amqpConn,
	}
	for _, o := range option {
		o(s)
	}
	return s
}

func (srv *RabbitMQTransport) Serve() {
	amqpChannel, err := srv.amqpConn.Channel()
	if err != nil {
		srv.logger.Error(fmt.Sprintf("failed to serve", zap.Error(err)))
	}

	// Register Repository , UseCase, Consumer
	newEventRepository := eventRepository.NewEventRepository(srv.mongoDBInfrastructure, srv.logger)
	newEventUseCase := eventUseCase.NewEventStoreUseCase(newEventRepository, srv.logger)
	newEventConsumer := consumer.NewEventConsumer(amqpChannel, newEventUseCase, srv.logger)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		log.Println("RabbitMQ shutdown...")
	}()

	go func() {
		defer cancel()
		if err = newEventConsumer.EventCreated(ctx); err != nil {
			srv.logger.Error(fmt.Sprintf("failed to ProductCreated : %s", zap.Error(err).String))
		}
	}()

	<-ctx.Done()

	if err = newEventConsumer.Close(ctx); err != nil {
		srv.logger.Error(fmt.Sprintf("failed to close a connection: %v", zap.Error(err)))
	}
}
