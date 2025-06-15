package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/enum"
	mailHogInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/mailhog"
	eventRpc "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/event/v1"
	notificationRpc "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/notification/v1"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/util"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *notificationEmailUseCase) SendNotificationEmailOTP(ctx context.Context, requestId string, req *notificationRpc.SendOtpEmailNotificationRequest) error {
	var (
		err        error
		eventStore = &eventRpc.EventStore{
			RequestId:     requestId,
			Service:       config.Get().ServiceName,
			EventType:     config.Get().QueueNotificationEmailOtpCreated,
			Status:        config.Get().CommonSagaStatusSuccess,
			PreviousState: nil,
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		}
	)

	ctxTimeout, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.SendUserOtpEmailNotification")
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

		if err = u.rabbitmqInfrastructure.Publish(ctxTimeout, requestId, config.Get().ExchangeEvent, config.Get().QueueEventCreated, eventStoreMessage); err != nil {
			u.logger.Error(fmt.Sprintf("error creating product event store: %s", err.Error()))
			return
		}
	}(err, eventStore)

	notificationType, err := enum.NotificationTypeParseIntToNotificationType(int(req.NotificationType))
	if err != nil {
		u.logger.Error(fmt.Sprintf("error parsing email type: %s", err.Error()))
		return err
	}

	fetchTemplate, err := u.notificationRepository.FindNotificationTemplateByNotificationType(ctxTimeout, notificationType)
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
