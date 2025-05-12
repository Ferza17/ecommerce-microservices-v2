package rabbitmq

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/bootstrap"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	productConsumer "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/consumer"
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
		amqpConn   *amqp091.Connection
		logger     pkg.IZapLogger
		dependency *bootstrap.Bootstrap
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

	return &Server{
		amqpConn:   amqpConn,
		dependency: dependency,
		logger:     dependency.Logger,
	}
}

func (srv *Server) Serve() {
	amqpChannel, err := srv.amqpConn.Channel()
	if err != nil {
		srv.dependency.Logger.Error(fmt.Sprintf("failed to serve", zap.Error(err)))
	}

	productConsumer := productConsumer.NewProductConsumer(amqpChannel, srv.dependency.ProductUseCase, srv.dependency.Logger)

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
		if err = productConsumer.ProductCreated(ctx); err != nil {
			srv.dependency.Logger.Error(fmt.Sprintf("failed to ProductCreated : %s", zap.Error(err).String))
		}
	}()

	go func() {
		if err = productConsumer.ProductUpdated(ctx); err != nil {
			srv.dependency.Logger.Error(fmt.Sprintf("failed to ProductUpdated : %s", zap.Error(err).String))
		}
	}()

	<-ctx.Done()

	if err = amqpChannel.Close(); err != nil {
		srv.logger.Error(fmt.Sprintf("failed to close channel %v", zap.Error(err)))
		return
	}
}
