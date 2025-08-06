package presenter

import (
	"context"
	"fmt"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (p *AuthPresenter) AuthUserRegister(ctx context.Context, req *pb.AuthUserRegisterRequest) (*emptypb.Empty, error) {
	ctx, span := p.telemetryInfrastructure.StartSpanFromContext(ctx, "AuthPresenter.AuthUserRegister")
	defer span.End()
	requestID := pkgContext.GetRequestIDFromContext(ctx)

	if err := req.Validate(); err != nil {
		p.logger.Error("AuthPresenter.AuthUserRegister", zap.String("requestID", requestID), zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	workflowRun, err := p.temporal.
		StartWorkflow(ctx, requestID, p.authWorkflow.AuthUserRegisterWorkflow, requestID, req)
	if err != nil {
		p.logger.Error(fmt.Sprintf("failed to start workflow : %v", zap.Error(err)))
	}

	var resp emptypb.Empty
	if err = workflowRun.Get(ctx, &resp); err != nil {
		p.logger.Error("AuthUserRegisterWorkflow failed", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &resp, nil
}
