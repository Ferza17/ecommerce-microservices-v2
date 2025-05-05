package rabbitmq

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/postgresql"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/rabbitmq"
	productConsumer "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/consumer"
	productpgRepo "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/repository/postgresql"
	productUseCase "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg"
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

		logger                 pkg.IZapLogger
		postgresqlConnector    postgresql.IPostgreSQLInfrastructure
		rabbitmqInfrastructure rabbitmq.IRabbitMQInfrastructure
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

	// Register Repository & UseCase
	newProductPgRepo := productpgRepo.NewProductPostgresqlRepository(srv.postgresqlConnector, srv.logger)
	newProductUseCase := productUseCase.NewProductUseCase(newProductPgRepo, srv.rabbitmqInfrastructure, srv.logger)

	newProductConsumer := productConsumer.NewProductConsumer(amqpChannel, newProductUseCase, srv.logger)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		log.Println("AMQP shutdown...")
	}()

	go func() {
		defer cancel()
		if err = newProductConsumer.ProductCreated(ctx); err != nil {
			srv.logger.Error(fmt.Sprintf("failed to ProductCreated : %s", zap.Error(err).String))
		}
	}()

	go func() {
		if err = newProductConsumer.ProductUpdated(ctx); err != nil {
			srv.logger.Error(fmt.Sprintf("failed to ProductUpdated : %s", zap.Error(err).String))
		}
	}()

	<-ctx.Done()
}
