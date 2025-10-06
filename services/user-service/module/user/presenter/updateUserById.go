package presenter

import (
	"context"

	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (p *UserPresenter) UpdateUserById(ctx context.Context, req *pb.UpdateUserByIdRequest) (*emptypb.Empty, error) {
	var (
		err       error
		requestID = pkgContext.GetRequestIDFromContext(ctx)
	)

	ctx, span := p.telemetryInfrastructure.StartSpanFromContext(ctx, "UserPresenter.UpdateUserById")
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()

	if err = req.Validate(); err != nil {
		p.logger.Error("UserPresenter.UpdateUserById", zap.String("requestID", requestID), zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if _, err = p.userUseCase.UpdateUserById(ctx, requestID, req); err != nil {
		p.logger.Error("UserPresenter.UpdateUserById", zap.String("requestID", requestID), zap.Error(err))
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
