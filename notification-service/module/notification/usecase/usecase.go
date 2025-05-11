package usecase

import (
	"context"
	mailHogInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/mailhog"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/pb"
	notificationRepository "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/notification/repository/mongodb"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg"
)

type (
	INotificationUseCase interface {
		SendLoginEmailNotification(ctx context.Context, requestId string, req *pb.SendLoginEmailNotificationRequest) (*pb.SendLoginEmailNotificationResponse, error)
		SendUserVerificationEmailNotification(ctx context.Context, requestId string, req *pb.SendUserVerificationEmailNotificationRequest) error
	}

	notificationUseCase struct {
		notificationRepository notificationRepository.INotificationRepository
		rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure
		mailHogInfrastructure  mailHogInfrastructure.IMailhogInfrastructure
		logger                 pkg.IZapLogger
	}
)

func NewEventStoreUseCase(
	notificationRepository notificationRepository.INotificationRepository,
	rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	mailHogInfrastructure mailHogInfrastructure.IMailhogInfrastructure,
	logger pkg.IZapLogger) INotificationUseCase {
	return &notificationUseCase{
		notificationRepository: notificationRepository,
		rabbitmqInfrastructure: rabbitmqInfrastructure,
		mailHogInfrastructure:  mailHogInfrastructure,
		logger:                 logger,
	}
}
