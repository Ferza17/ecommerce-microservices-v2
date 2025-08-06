package workflow

import (
	"fmt"
	pb "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/v1/notification"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/zap"
	"time"
)

func (w *emailWorkflow) SendNotificationEmailPaymentOrderCreatedWorkflow(wCtx workflow.Context, requestId string, req *pb.SendEmailPaymentOrderCreateRequest) error {
	logger := workflow.GetLogger(wCtx)
	activityOptions := workflow.ActivityOptions{
		WaitForCancellation: true,
		StartToCloseTimeout: time.Minute * 1,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Millisecond * 60,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Minute * 1,
			MaximumAttempts:    1,
		},
	}
	var err error
	if err = workflow.
		ExecuteActivity(workflow.WithActivityOptions(wCtx, activityOptions), w.notificationUseCase.SendNotificationEmailPaymentOrderCreated, requestId, req).
		Get(wCtx, nil); err != nil {
		logger.Error(fmt.Sprintf("failed to execute workflow : %v", zap.Error(err)))
		return err
	}

	return nil
}
