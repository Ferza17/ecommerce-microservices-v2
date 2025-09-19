package usecase

import (
	"context"
	kafkaInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/kafka"
	mailHogInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/mailhog"
	paymentService "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/services/payment"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/telemetry"
	notificationRpc "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/v1/notification"
	notificationRepository "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/email/repository/mongodb"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	"github.com/google/wire"
)

type (
	INotificationEmailUseCase interface {
		SendNotificationEmailOTP(ctx context.Context, requestId string, req *notificationRpc.SendOtpEmailNotificationRequest) error
		SendNotificationEmailPaymentOrderCreated(ctx context.Context, requestId string, req *notificationRpc.SendEmailPaymentOrderCreateRequest) error
	}

	notificationEmailUseCase struct {
		notificationRepository  notificationRepository.INotificationEmailRepository
		kafkaInfrastructure     kafkaInfrastructure.IKafkaInfrastructure
		mailHogInfrastructure   mailHogInfrastructure.IMailhogInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		paymentSvc              paymentService.IPaymentService
		logger                  logger.IZapLogger
	}
)

var Set = wire.NewSet(NewNotificationEmailUseCase)

func NewNotificationEmailUseCase(
	notificationRepository notificationRepository.INotificationEmailRepository,
	kafkaInfrastructure kafkaInfrastructure.IKafkaInfrastructure,
	mailHogInfrastructure mailHogInfrastructure.IMailhogInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	paymentSvc paymentService.IPaymentService,
	logger logger.IZapLogger) INotificationEmailUseCase {
	c := &notificationEmailUseCase{
		notificationRepository:  notificationRepository,
		kafkaInfrastructure:     kafkaInfrastructure,
		mailHogInfrastructure:   mailHogInfrastructure,
		telemetryInfrastructure: telemetryInfrastructure,
		paymentSvc:              paymentSvc,
		logger:                  logger,
	}
	return c
}
