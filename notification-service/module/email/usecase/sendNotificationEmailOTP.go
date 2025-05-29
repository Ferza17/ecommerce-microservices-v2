package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/enum"
	mailHogInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/mailhog"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/pb"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/util"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *notificationEmailUseCase) SendNotificationEmailOTP(ctx context.Context, requestId string, req *pb.SendOtpEmailNotificationRequest) error {
	var (
		err        error
		eventStore = &pb.EventStore{
			RequestId:     requestId,
			Service:       enum.NotificationService.String(),
			EventType:     enum.NOTIFICATION_EMAIL_OTP.String(),
			Status:        enum.SUCCESS.String(),
			PreviousState: nil,
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		}
	)

	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.SendUserOtpEmailNotification")
	defer func(err error, eventStore *pb.EventStore) {
		defer span.End()
		if err != nil {
			eventStore.Status = enum.FAILED.String()
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

		if err = u.rabbitmqInfrastructure.Publish(ctx, requestId, enum.EventExchange, enum.EVENT_CREATED, eventStoreMessage); err != nil {
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
