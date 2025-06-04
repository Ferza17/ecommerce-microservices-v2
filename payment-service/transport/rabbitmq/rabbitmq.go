package rabbitmq

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/consumer"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type (
	IRabbitMQServer interface {
		Serve()
	}

	rabbitMQServer struct {
		rabbitMQ rabbitmq.IRabbitMQInfrastructure
		logger   logger.IZapLogger
	}
)

func NewRabbitMQServer(rabbitMQ rabbitmq.IRabbitMQInfrastructure, logger logger.IZapLogger) IRabbitMQServer {
	return &rabbitMQServer{
		rabbitMQ: rabbitMQ,
		logger:   logger,
	}
}

func (r rabbitMQServer) Serve() {
	ctx, cancel := context.WithCancel(context.Background())

	paymentConsumer := consumer.ProvidePaymentConsumer()
	go func() {
		defer cancel()
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		log.Println("AMQP shutdown...")
	}()

	go func() {
		if err := paymentConsumer.PaymentOrderCreated(ctx); err != nil {
			r.logger.Error(fmt.Sprintf("Err PaymentOrderCreated : %v", err))
		}
	}()

	go func() {
		if err := paymentConsumer.PaymentOrderDelayedCancelled(ctx); err != nil {
			r.logger.Error(fmt.Sprintf("Err PaymentOrderDelayedCancelled : %v", err))
		}
	}()

	<-ctx.Done()
}
