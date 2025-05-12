package rabbitmq

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/bootstrap"
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
		amqpConn        *amqp091.Connection
		logger          pkg.IZapLogger
		dependency      *bootstrap.Bootstrap
		productConsumer productConsumer.IProductConsumer
	}
)

func NewServer(dependency *bootstrap.Bootstrap) *Server {
	return &Server{
		dependency:      dependency,
		logger:          dependency.Logger,
		productConsumer: dependency.ProductConsumer,
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
		defer cancel()
		if err := srv.productConsumer.ProductCreated(ctx); err != nil {
			srv.dependency.Logger.Error(fmt.Sprintf("failed to ProductCreated : %s", zap.Error(err).String))
		}
	}()

	go func() {
		if err := srv.productConsumer.ProductUpdated(ctx); err != nil {
			srv.dependency.Logger.Error(fmt.Sprintf("failed to ProductUpdated : %s", zap.Error(err).String))
		}
	}()

	<-ctx.Done()
}
