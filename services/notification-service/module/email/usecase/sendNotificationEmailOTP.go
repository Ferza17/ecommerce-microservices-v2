package usecase

import (
	"context"
	"errors"
	"fmt"
	mailHogInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/mailhog"
	notificationRpc "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/v1/notification"
)

func (u *notificationEmailUseCase) SendNotificationEmailOTP(ctx context.Context, requestId string, req *notificationRpc.SendOtpEmailNotificationRequest) error {
	var (
		err error
	)

	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "NotificationUseCase.SendUserOtpEmailNotification")
	defer span.End()

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
