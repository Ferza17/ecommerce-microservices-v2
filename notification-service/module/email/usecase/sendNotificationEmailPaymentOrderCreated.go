package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/enum"
	mailHogInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/mailhog"
	notificationRpc "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/v1/notification"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *notificationEmailUseCase) SendNotificationEmailPaymentOrderCreated(ctx context.Context, requestId string, req *notificationRpc.SendEmailPaymentOrderCreateRequest) error {
	var (
		err error
	)

	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "UseCase.SendUserOtpEmailNotification")
	defer span.End()

	notificationType, err := enum.NotificationTypeParseIntToNotificationType(int(req.NotificationType))
	if err != nil {
		u.logger.Error(fmt.Sprintf("error parsing email type: %s", err.Error()))
		return err
	}

	fetchTemplate, err := u.notificationRepository.FindNotificationTemplateByNotificationType(ctx, notificationType)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error finding email template by email type: %s", err.Error()))
		return status.Error(codes.Internal, "error finding email template by email type")
	}

	if fetchTemplate == nil {
		u.logger.Error(fmt.Sprintf("email template not found"))
		return status.Error(codes.NotFound, "email template not found")
	}

	var (
		templateVars = map[string]any{
			"Code":   req.Payment.Code,
			"Status": req.Payment.Status.String(),
			//"Provider":     req.Payment.Provider,
			"CreatedAt":  req.Payment.CreatedAt,
			"TotalPrice": req.Payment.TotalPrice,
			//"PaymentItems": req.Payment.Items,
		}
	)

	// TODO: Assign request to template vars
	if err = u.mailHogInfrastructure.SendMail(&mailHogInfrastructure.Mailer{
		Subject:      "🤯 PAYMENT ORDER CREATED 🤯",
		To:           req.Email,
		Template:     fetchTemplate.Template,
		TemplateVars: templateVars,
	}); err != nil {
		return err
	}

	return nil
}
