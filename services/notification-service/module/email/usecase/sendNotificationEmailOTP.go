package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/kafka"
	mailHogInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/mailhog"
	pbEvent "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/v1/event"
	pbNotification "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/v1/notification"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/context"
)

func (u *notificationEmailUseCase) SendNotificationEmailOTP(ctx context.Context, requestId string, req *pbNotification.SendOtpEmailNotificationRequest) error {
	var (
		err error
	)

	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "NotificationUseCase.SendUserOtpEmailNotification")

	defer func(err error) {
		reserveEvent := &pbEvent.ReserveEvent{
			SagaId:        pkgContext.GetRequestIDFromContext(ctx),
			AggregateType: "users",
		}

		if req.NotificationType == pbNotification.NotificationTypeEnum_NOTIFICATION_EMAIL_USER_REGISTER_OTP && err == nil {
			if err = u.kafkaInfrastructure.PublishWithSchema(ctx, config.Get().BrokerKafkaTopicUsers.ConfirmUserUserCreated, reserveEvent.SagaId, kafka.PROTOBUF_SCHEMA, reserveEvent); err != nil {
				u.logger.Error(fmt.Sprintf("failed to publish message to topic %s: %v", config.Get().BrokerKafkaTopicUsers.ConfirmUserUserCreated, err))
			}
		}

		if req.NotificationType == pbNotification.NotificationTypeEnum_NOTIFICATION_EMAIL_USER_REGISTER_OTP && err != nil {
			if err = u.kafkaInfrastructure.PublishWithSchema(ctx, config.Get().BrokerKafkaTopicUsers.CompensateUserUserCreated, reserveEvent.SagaId, kafka.PROTOBUF_SCHEMA, reserveEvent); err != nil {
				u.logger.Error(fmt.Sprintf("failed to publish message to topic confirm-snapshot-users-user_created: %v", err))
			}
		}

		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}(err)

	fetchTemplate, err := u.notificationRepository.FindNotificationTemplateByNotificationType(ctx, requestId, req.NotificationType)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error finding email template by email type: %s", err.Error()))
		return err
	}

	if fetchTemplate == nil {
		u.logger.Error(fmt.Sprintf("email template not found"))
		return errors.New("email template not found")
	}

	if err = u.mailHogInfrastructure.SendMail(&mailHogInfrastructure.Mailer{
		Subject:  "🤯 OTP VERIFICATION 🤯",
		To:       req.Email,
		Template: fetchTemplate.Template,
		TemplateVars: map[string]any{
			"otp": req.Otp,
		},
	}); err != nil {
		return err
	}

	return nil
}
