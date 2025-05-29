package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/enum"
	mailHogInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/mailhog"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/pb"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *notificationUseCase) SendLoginEmailNotification(ctx context.Context, requestId string, req *pb.SendLoginEmailNotificationRequest) (*pb.SendLoginEmailNotificationResponse, error) {
	var (
		err        error
		eventStore = &pb.EventStore{
			RequestId:     requestId,
			Service:       enum.NotificationService.String(),
			EventType:     enum.NOTIFICATION_LOGIN_CREATED.String(),
			Status:        enum.SUCCESS.String(),
			PreviousState: nil,
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		}
	)
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.SendLoginEmailNotification")

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
		u.logger.Error(fmt.Sprintf("error parsing notification type: %s", err.Error()))
		return nil, status.Error(codes.Internal, err.Error())
	}

	fetchTemplate, err := u.notificationRepository.FindNotificationTemplateByNotificationType(ctx, notificationType)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error finding notification template by notification type: %s", err.Error()))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if fetchTemplate == nil {
		u.logger.Error(fmt.Sprintf("notification template not found"))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	if err = u.mailHogInfrastructure.SendMail(&mailHogInfrastructure.Mailer{
		Subject:  "🤯 LOGIN 🤯",
		To:       req.Email,
		Template: fetchTemplate.Template,
		TemplateVars: map[string]any{
			"access_token":  req.AccessToken,
			"refresh_token": req.RefreshToken,
			"username":      req.Username,
		},
	}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return nil, nil
}
