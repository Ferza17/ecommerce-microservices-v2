package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/enum"
	mailHogInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/mailhog"
	notificationRpc "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/v1/notification"
	pb "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/v1/payment"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"time"
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

	fetchPayment, err := u.paymentSvc.FindPaymentById(ctx, requestId, &pb.FindPaymentByIdRequest{
		Id: req.PaymentId,
	})
	if err != nil {
		u.logger.Error(fmt.Sprintf("error finding email payment by id: %s", err.Error()))
		return status.Error(codes.Internal, "error finding email payment by id")
	}

	var (
		templateVars = map[string]any{
			"Code":   fetchPayment.Data.Payment.Code,
			"Status": strings.ToLower(fetchPayment.Data.Payment.Status.String()),
			"Provider": struct {
				Name      string
				Method    string
				CreatedAt time.Time
				UpdatedAt time.Time
			}{
				Name:      fetchPayment.Data.Provider.Name,
				Method:    fetchPayment.Data.Provider.Method.String(),
				CreatedAt: fetchPayment.Data.Provider.CreatedAt.AsTime(),
				UpdatedAt: fetchPayment.Data.Provider.UpdatedAt.AsTime(),
			},
			"CreatedAt":    fetchPayment.Data.Payment.CreatedAt.AsTime(),
			"TotalPrice":   fetchPayment.Data.Payment.TotalPrice,
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

	for _, item := range fetchPayment.Data.PaymentItems {
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
