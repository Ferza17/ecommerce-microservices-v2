package rabbitmq

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/bootstrap"
	notificationConsumer "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/notification/consumer"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type (
	RabbitMQTransport struct {
		logger               pkg.IZapLogger
		notificationConsumer notificationConsumer.INotificationConsumer
	}
)

func NewServer(dependency *bootstrap.Bootstrap) *RabbitMQTransport {
	return &RabbitMQTransport{
		logger:               dependency.Logger,
		notificationConsumer: dependency.NotificationConsumer,
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
		if err := srv.notificationConsumer.NotificationLoginCreated(ctx); err != nil {
			srv.logger.Error(fmt.Sprintf("failed to ProductCreated : %s", zap.Error(err).String))
		}
	}()

	go func() {
		defer cancel()
		if err := srv.notificationConsumer.NotificationUserCreated(ctx); err != nil {
			srv.logger.Error(fmt.Sprintf("failed to ProductCreated : %s", zap.Error(err).String))
		}
	}()

	<-ctx.Done()
}
