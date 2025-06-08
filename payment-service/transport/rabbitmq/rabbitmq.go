package rabbitmq

import (
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/rabbitmq"
	paymentConsumer "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/consumer"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IRabbitMQServer interface {
		Serve()
	}

	rabbitMQServer struct {
		rabbitMQ rabbitmq.IRabbitMQInfrastructure

		paymentConsumer paymentConsumer.IPaymentConsumer

		logger logger.IZapLogger
	}
)

// NewRabbitMQServer creates and returns a new instance of RabbitMQServer with all dependencies.
func NewRabbitMQServer(
	rabbitMQ rabbitmq.IRabbitMQInfrastructure,
	paymentConsumer paymentConsumer.IPaymentConsumer,
	logger logger.IZapLogger,
) IRabbitMQServer {
	return &rabbitMQServer{
		rabbitMQ:        rabbitMQ,
		paymentConsumer: paymentConsumer,
		logger:          logger,
	}
}

// Set is a Wire provider set for rabbitMQServer dependencies.
var Set = wire.NewSet(
	NewRabbitMQServer,
)

func (r rabbitMQServer) Serve() {
	//ctx, cancel := context.WithCancel(context.Background())

	//paymentConsumer := consumer.ProvidePaymentConsumer()
	//go func() {
	//	defer cancel()
	//	sigChan := make(chan os.Signal, 1)
	//	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	//	<-sigChan
	//	log.Println("AMQP shutdown...")
	//}()
	//
	//go func() {
	//	if err := paymentConsumer.PaymentOrderCreated(ctx); err != nil {
	//		r.logger.Error(fmt.Sprintf("Err PaymentOrderCreated : %v", err))
	//	}
	//}()
	//
	//go func() {
	//	if err := paymentConsumer.PaymentOrderDelayedCancelled(ctx); err != nil {
	//		r.logger.Error(fmt.Sprintf("Err PaymentOrderDelayedCancelled : %v", err))
	//	}
	//}()

	//<-ctx.Done()
}
