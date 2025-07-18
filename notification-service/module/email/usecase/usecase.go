package usecase

import (
	"context"
	mailHogInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/mailhog"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/rabbitmq"
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
		rabbitmqInfrastructure  rabbitmqInfrastructure.IRabbitMQInfrastructure
		mailHogInfrastructure   mailHogInfrastructure.IMailhogInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  logger.IZapLogger
	}
)

var Set = wire.NewSet(NewEventStoreUseCase)

func NewEventStoreUseCase(
	notificationRepository notificationRepository.INotificationEmailRepository,
	rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	mailHogInfrastructure mailHogInfrastructure.IMailhogInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger) INotificationEmailUseCase {
	return &notificationEmailUseCase{
		notificationRepository:  notificationRepository,
		rabbitmqInfrastructure:  rabbitmqInfrastructure,
		mailHogInfrastructure:   mailHogInfrastructure,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
