package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/enum"
	mailHogInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/mailhog"
	eventRpc "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/event/v1"
	notificationRpc "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/notification/v1"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *notificationEmailUseCase) SendNotificationEmailPaymentOrderCreated(ctx context.Context, requestId string, req *notificationRpc.SendEmailPaymentOrderCreateRequest) error {
	var (
		err        error
		eventStore = &eventRpc.EventStore{
			RequestId:     requestId,
			Service:       config.Get().ServiceName,
			EventType:     config.Get().QueueNotificationEmailPaymentOrderCreated,
			Status:        config.Get().CommonSagaStatusSuccess,
			PreviousState: nil,
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		}
	)

	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.SendUserOtpEmailNotification")
	defer func(err error, eventStore *eventRpc.EventStore) {
		defer span.End()
		if err != nil {
			eventStore.Status = config.Get().CommonSagaStatusFailed
		}

		payload, err := util.ConvertStructToProtoStruct(req)
		if err != nil {
			u.logger.Error(fmt.Sprintf("error converting struct to proto struct: %s", err.Error()))
		}
		eventStore.Payload = payload

		eventStoreMessage, err := proto.Marshal(eventStore)
		if err != nil {
			u.logger.Error(fmt.Sprintf("error marshaling message: %s", err.Error()))
		}

		if err = u.rabbitmqInfrastructure.Publish(ctx, requestId, config.Get().ExchangeEvent, config.Get().QueueEventCreated, eventStoreMessage); err != nil {
			u.logger.Error(fmt.Sprintf("error creating product event store: %s", err.Error()))
			return
		}
	}(err, eventStore)

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
			"Code":         req.Payment.Code,
			"Status":       req.Payment.Status.String(),
			"Provider":     req.Payment.Provider,
			"CreatedAt":    req.Payment.CreatedAt,
			"TotalPrice":   req.Payment.TotalPrice,
			"PaymentItems": req.Payment.Items,
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
