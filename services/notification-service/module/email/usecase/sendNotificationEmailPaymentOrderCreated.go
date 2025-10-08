package usecase

import (
	"context"
	"fmt"
	"strings"
	"time"

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
	defer func(err error) {
		if err == nil {
			// TODO:
			// Publish to topic confirm-snapshot-payments-payment_order_created

		} else {
			// TODO:
			// Publish to topic compensate-snapshot-payments-payment_order_created

			span.RecordError(err)
		}
		span.End()
	}(err)

	fetchTemplate, err := u.notificationRepository.FindNotificationTemplateByNotificationType(ctx, requestId, req.NotificationType)
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
			"Status": strings.ToLower(req.Payment.Status.String()),
			"Provider": struct {
				Name      string
				Method    string
				CreatedAt time.Time
				UpdatedAt time.Time
			}{
				Name:      req.PaymentProvider.Name,
				Method:    req.PaymentProvider.Method.String(),
				CreatedAt: req.PaymentProvider.CreatedAt.AsTime(),
				UpdatedAt: req.PaymentProvider.UpdatedAt.AsTime(),
			},
			"CreatedAt":    req.Payment.CreatedAt.AsTime(),
			"TotalPrice":   req.Payment.TotalPrice,
			"PaymentItems": "",
		}
	)

	var paymentItems []struct {
		ProductID string
		Amount    float64
		Qty       uint32
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	for _, item := range req.Payment.Items {
		paymentItems = append(paymentItems, struct {
			ProductID string
			Amount    float64
			Qty       uint32
			CreatedAt time.Time
			UpdatedAt time.Time
		}{ProductID: item.ProductId, Amount: item.Amount, Qty: uint32(item.Qty), CreatedAt: item.CreatedAt.AsTime(), UpdatedAt: item.UpdatedAt.AsTime()})
	}

	templateVars["PaymentItems"] = paymentItems

	if err = u.mailHogInfrastructure.SendMail(&mailHogInfrastructure.Mailer{
		Subject:      "🤯 PAYMENT ORDER CREATED 🤯",
		To:           req.Email,
		Template:     fetchTemplate.Template,
		TemplateVars: templateVars,
	}); err != nil {
		u.logger.Error(fmt.Sprintf("error sending email template: %s", err.Error()))
		return err
	}
	return nil
}
