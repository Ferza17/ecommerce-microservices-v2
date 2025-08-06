package workflow

import (
	"fmt"
	pb "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/v1/notification"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/zap"
	"time"
)

func (w *emailWorkflow) NotificationEmailOTPWorkflow(wCtx workflow.Context, requestId string, req *pb.SendOtpEmailNotificationRequest) error {
	logger := workflow.GetLogger(wCtx)
	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 1,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Millisecond * 60,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Minute * 3,
			MaximumAttempts:    1,
		},
	}

	var err error
	if err = workflow.
		ExecuteActivity(workflow.WithActivityOptions(wCtx, activityOptions), w.notificationUseCase.SendNotificationEmailOTP, requestId, req).
		Get(wCtx, &err); err != nil {
		logger.Error(fmt.Sprintf("failed to execute workflow : %v", zap.Error(err)))
		return err
	}

	return nil
}
