package rabbitmq

import (
	"context"
	"fmt"
	notificationEmailConsumer "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/email/consumer"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	"github.com/google/wire"
	"go.uber.org/zap"
	"sync"
)

type (
	RabbitMQTransport struct {
		logger                    logger.IZapLogger
		notificationEmailConsumer notificationEmailConsumer.INotificationEmailConsumer
	}
)

var Set = wire.NewSet(NewServer)

func NewServer(
	logger logger.IZapLogger,
	notificationEmailConsumer notificationEmailConsumer.INotificationEmailConsumer,
) *RabbitMQTransport {
	return &RabbitMQTransport{
		logger:                    logger,
		notificationEmailConsumer: notificationEmailConsumer,
	}
}

func (srv *RabbitMQTransport) Serve(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer cancel()
		if err := srv.notificationEmailConsumer.NotificationEmailOTP(ctx); err != nil {
			srv.logger.Error(fmt.Sprintf("failed to ProductCreated : %s", zap.Error(err).String))
		}
	}()

	wg.Wait()
	cancel()
}
