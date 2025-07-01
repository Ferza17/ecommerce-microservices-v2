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

func (p *AuthPresenter) AuthServiceVerifyIsExcluded(ctx context.Context, req *pb.AuthServiceVerifyIsExcludedRequest) (*pb.AuthServiceVerifyIsExcludedResponse, error) {
	ctx, span := p.telemetryInfrastructure.Tracer(ctx, "AuthPresenter.AuthServiceVerifyIsExcluded")
	defer span.End()
	requestID := pkgContext.GetRequestIDFromContext(ctx)

	if err := p.customValidationAuthServiceVerifyIsExcludedRequest(requestID, req); err != nil {
		p.logger.Error("AuthPresenter.AuthServiceVerifyIsExcluded", zap.String("requestId", requestID), zap.Error(err))
		return nil, err
	}

	resp, err := p.authUseCase.AuthServiceVerifyIsExcluded(ctx, requestID, req)
	if err != nil {
		p.logger.Error("AuthPresenter.AuthServiceVerifyIsExcluded", zap.String("requestId", requestID), zap.Error(err))
		return nil, err
	}

	return resp, nil
}

func (p *AuthPresenter) customValidationAuthServiceVerifyIsExcludedRequest(requestId string, req *pb.AuthServiceVerifyIsExcludedRequest) error {
	errCounter := 0
	if req.FullMethodName == nil {
		errCounter += 1
	}

	if req.HttpUrl == nil {
		errCounter += 1
	}

	if req.HttpMethod != nil {
		if req.HttpMethod == nil {
			p.logger.Error("AuthPresenter.customValidationAuthServiceVerifyIsExcludedRequest", zap.String("requestID", requestId), zap.Error(errors.New("http method is required")))
			return status.Error(codes.InvalidArgument, "http method is required")
		}
	}

	if errCounter > 1 {
		p.logger.Error("AuthPresenter.customValidationAuthServiceVerifyIsExcludedRequest", zap.String("requestID", requestId), zap.Error(errors.New("please provide at least one valid http_url or full_method_name (for RPC call)")))
		return status.Error(codes.FailedPrecondition, "please provide at least one valid http_url or full_method_name (for RPC call)")
	}

	return nil
}
