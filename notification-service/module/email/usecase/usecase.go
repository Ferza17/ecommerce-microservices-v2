package usecase

import (
	"context"
	mailHogInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/mailhog"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/pb"
	notificationRepository "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/email/repository/mongodb"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg"
)

type (
	INotificationEmailUseCase interface {
		SendNotificationEmailOTP(ctx context.Context, requestId string, req *pb.SendOtpEmailNotificationRequest) error
	}

	notificationEmailUseCase struct {
		notificationRepository  notificationRepository.INotificationEmailRepository
		rabbitmqInfrastructure  rabbitmqInfrastructure.IRabbitMQInfrastructure
		mailHogInfrastructure   mailHogInfrastructure.IMailhogInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  pkg.IZapLogger
	}
)

func NewEventStoreUseCase(
	notificationRepository notificationRepository.INotificationEmailRepository,
	rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	mailHogInfrastructure mailHogInfrastructure.IMailhogInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkg.IZapLogger) INotificationEmailUseCase {
	return &notificationEmailUseCase{
		notificationRepository:  notificationRepository,
		rabbitmqInfrastructure:  rabbitmqInfrastructure,
		mailHogInfrastructure:   mailHogInfrastructure,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
