package usecase

import (
	"context"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
)

func (u *authUseCase) AuthServiceVerifyIsExcluded(ctx context.Context, requestId string, req *pb.AuthServiceVerifyIsExcludedRequest) (*pb.AuthServiceVerifyIsExcludedResponse, error) {
	var (
		isExcluded = false
		err        error
	)

	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "AuthUseCase.AuthServiceVerifyIsExcluded")
	defer span.End()

	if req.FullMethodName != nil {
		isExcluded, err = u.accessControlUseCase.IsExcludedRPC(ctx, requestId, *req.FullMethodName)
		if err != nil {
			u.logger.Error("AuthUseCase.AuthServiceVerifyIsExcluded", zap.String("requestId", requestId), zap.Error(err))
			return nil, err
		}
	}

	if req.HttpMethod != nil && req.HttpUrl != nil {
		isExcluded, err = u.accessControlUseCase.IsExcludedHTTP(ctx, requestId, *req.HttpMethod, *req.HttpUrl)
		if err != nil {
			u.logger.Error("AuthUseCase.AuthServiceVerifyIsExcluded", zap.String("requestId", requestId), zap.Error(err))
		}
	}

	return &pb.AuthServiceVerifyIsExcludedResponse{
		Error:   "",
		Message: codes.OK.String(),
		Code:    uint32(codes.OK),
		Data: &pb.AuthServiceVerifyIsExcludedResponse_AuthServiceVerifyIsExcludedResponseData{
			IsExcluded: isExcluded,
		},
	}, err
}
