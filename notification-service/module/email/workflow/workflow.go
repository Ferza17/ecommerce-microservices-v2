package workflow

import (
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/temporal"
	pb "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/v1/notification"
	notificationUseCase "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/email/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	"github.com/google/wire"
	"go.temporal.io/sdk/workflow"
)

type (
	IEmailWorkflow interface {
		NotificationEmailOTPWorkflow(wCtx workflow.Context, requestId string, req *pb.SendOtpEmailNotificationRequest) error
		SendNotificationEmailPaymentOrderCreatedWorkflow(wCtx workflow.Context, requestId string, req *pb.SendEmailPaymentOrderCreateRequest) error
	}

	emailWorkflow struct {
		notificationUseCase notificationUseCase.INotificationEmailUseCase
		temporal            temporal.ITemporalInfrastructure
		logger              logger.IZapLogger
	}
)

var Set = wire.NewSet(NewEmailWorkflow)

func NewEmailWorkflow(
	notificationUseCase notificationUseCase.INotificationEmailUseCase,
	temporal temporal.ITemporalInfrastructure,
	logger logger.IZapLogger,
) IEmailWorkflow {
	c := &emailWorkflow{
		notificationUseCase: notificationUseCase,
		temporal:            temporal,
		logger:              logger,
	}
	c.temporal = c.temporal.
		RegisterWorkflow(c.NotificationEmailOTPWorkflow).
		RegisterWorkflow(c.SendNotificationEmailPaymentOrderCreatedWorkflow)

	return c
}
