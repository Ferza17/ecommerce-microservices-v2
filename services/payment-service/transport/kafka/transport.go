package kafka

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	kafkaInfrastructure "github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/kafka"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	paymentConsumer "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/consumer"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	pkgWorker "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/worker"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/attribute"
	"strings"
	"time"
)

type (
	Transport struct {
		workerPool              *pkgWorker.WorkerPool
		paymentConsumer         paymentConsumer.IPaymentConsumer
		kafkaInfrastructure     kafkaInfrastructure.IKafkaInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  logger.IZapLogger
		topics                  []string
	}
)

var Set = wire.NewSet(
	NewTransport,
)

func NewTransport(
	paymentConsumer paymentConsumer.IPaymentConsumer,
	kafkaInfrastructure kafkaInfrastructure.IKafkaInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger,
) *Transport {
	return &Transport{
		workerPool: pkgWorker.NewWorkerPoolKafkaTaskQueue(
			"Kafka Consumer", 9, 1000),
		paymentConsumer:         paymentConsumer,
		kafkaInfrastructure:     kafkaInfrastructure,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
		topics: []string{
			config.Get().BrokerKafkaTopicPayments.PaymentOrderCreated,
			config.Get().BrokerKafkaTopicPayments.PaymentOrderCreatedDelayed,
		},
	}
}

func (srv *Transport) Serve(mainCtx context.Context) error {
	srv.workerPool.Start()

	if err := srv.kafkaInfrastructure.SetupTopics(srv.topics); err != nil {
		srv.logger.Error(fmt.Sprintf("failed to setup kafka topics: %v", err))
		return err
	}

	run := true
	for run {
		select {
		case <-mainCtx.Done():
			srv.workerPool.Stop()
			run = false
		default:
			msg, err := srv.kafkaInfrastructure.ReadMessage(time.Second * 2)
			if err != nil {
				srv.logger.Error(fmt.Sprintf("failed to read message: %v", err))
				return err
			}

			if msg == nil {
				continue
			}

			var (
				requestId   string
				childCtx, _ = context.WithTimeout(mainCtx, 20*time.Second)
			)

			for _, header := range msg.Headers {
				if strings.ToLower(header.Key) == strings.ToLower(pkgContext.CtxKeyRequestID) {
					requestId = string(header.Value)
					childCtx = pkgContext.SetRequestIDToContext(childCtx, requestId)
				}

				if strings.ToLower(header.Key) == strings.ToLower(pkgContext.CtxKeyAuthorization) {
					childCtx = pkgContext.SetTokenAuthorizationToContext(childCtx, string(header.Value))
				}
			}
			childCtx, span := srv.telemetryInfrastructure.StartSpanFromKafkaHeader(childCtx, msg.Headers, "KafkaTransport")
			task := pkgWorker.KafkaTaskQueue{
				Message: msg,
				Ctx:     childCtx,
			}
			if msg.TopicPartition.Topic != nil {
				span.SetAttributes(attribute.String("messaging.destination", *msg.TopicPartition.Topic))
				span.SetAttributes(attribute.String(pkgContext.CtxKeyRequestID, requestId))

				switch *msg.TopicPartition.Topic {
				case config.Get().BrokerKafkaTopicPayments.PaymentOrderCreated:
					task.Handler = func(ctx context.Context, message *kafka.Message) error {
						return srv.paymentConsumer.SnapshotPaymentsPaymentOrderCreated(childCtx, message)
					}
				case config.Get().BrokerKafkaTopicPayments.PaymentOrderCreatedDelayed:
					task.Handler = func(ctx context.Context, message *kafka.Message) error {
						return srv.paymentConsumer.SnapshotPaymentsPaymentOrderCancelledDelayed(childCtx, message)
					}
				default:
					srv.logger.Error(fmt.Sprintf("invalid topic %s", *msg.TopicPartition.Topic))
					continue
				}
			}

			srv.workerPool.AddKafkaTaskQueue(task)
			span.End()
		}
	}
	return nil
}
