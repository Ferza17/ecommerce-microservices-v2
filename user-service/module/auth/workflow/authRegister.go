package workflow

import (
	"fmt"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

func (p *authWorkflow) AuthUserRegisterWorkflow(wCtx workflow.Context, requestId string, req *pb.AuthUserRegisterRequest) (*emptypb.Empty, error) {
	logger := workflow.GetLogger(wCtx)
	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 5,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Minute,
			MaximumAttempts:    3,
		},
	}
	var (
		resp emptypb.Empty
		err  error
	)
	
	if err = workflow.
		ExecuteActivity(workflow.WithActivityOptions(wCtx, activityOptions), p.authUseCase.AuthUserRegister, requestId, req).
		Get(wCtx, &resp); err != nil {
		logger.Error(fmt.Sprintf("failed to execute workflow : %v", zap.Error(err)))
		return nil, err
	}

	return &resp, nil
}
