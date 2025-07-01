package presenter

import (
	"context"
	"errors"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *AuthPresenter) AuthUserVerifyAccessControl(ctx context.Context, req *pb.AuthUserVerifyAccessControlRequest) (*pb.AuthUserVerifyAccessControlResponse, error) {
	ctx, span := p.telemetryInfrastructure.Tracer(ctx, "AuthPresenter.AuthUserVerifyAccessControlByToken")
	defer span.End()
	requestID := pkgContext.GetRequestIDFromContext(ctx)

	if err := req.Validate(); err != nil {
		p.logger.Error("AuthPresenter.AuthUserVerifyAccessControlByToken", zap.String("requestID", requestID), zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := p.customValidationAuthAccessControlByTokenRequest(requestID, req); err != nil {
		p.logger.Error("AuthPresenter.AuthUserVerifyAccessControlByToken", zap.String("requestID", requestID), zap.Error(err))
		return nil, err
	}

	resp, err := p.authUseCase.AuthUserVerifyAccessControl(ctx, requestID, req)
	if err != nil {
		p.logger.Error("AuthPresenter.AuthUserVerifyAccessControlByToken", zap.String("requestID", requestID), zap.Error(err))
		return nil, err
	}

	return resp, nil
}

func (p *AuthPresenter) customValidationAuthAccessControlByTokenRequest(requestId string, req *pb.AuthUserVerifyAccessControlRequest) error {
	errCounter := 0
	if req.FullMethodName == nil {
		errCounter += 1
	}

	if req.HttpUrl == nil {
		errCounter += 1
	}

	if req.HttpMethod != nil {
		if req.HttpMethod == nil {
			p.logger.Error("AuthPresenter.customValidationAuthAccessControlByTokenRequest", zap.String("requestID", requestId), zap.Error(errors.New("http method is required")))
			return status.Error(codes.InvalidArgument, "http method is required")
		}
	}

	if errCounter > 1 {
		p.logger.Error("AuthPresenter.customValidationAuthAccessControlByTokenRequest", zap.String("requestID", requestId), zap.Error(errors.New("please provide at least one valid http_url or full_method_name (for RPC call)")))
		return status.Error(codes.FailedPrecondition, "please provide at least one valid http_url or full_method_name (for RPC call)")
	}

	return nil
}
