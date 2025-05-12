package rabbitmq

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/bootstrap"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/module/event/consumer"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type (
	RabbitMQTransport struct {
		eventConsumer consumer.IEventConsumer
		logger        pkg.IZapLogger
	}
)

func NewServer(dependency *bootstrap.Bootstrap) *RabbitMQTransport {
	return &RabbitMQTransport{
		logger:        dependency.Logger,
		eventConsumer: dependency.EventConsumer,
	}
}

func (srv *RabbitMQTransport) Serve() {
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
		if err := srv.eventConsumer.EventCreated(ctx); err != nil {
			srv.logger.Error(fmt.Sprintf("failed to ProductCreated : %s", zap.Error(err).String))
		}
	}()

	<-ctx.Done()
}
