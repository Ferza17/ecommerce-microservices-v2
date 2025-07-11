package presenter

import (
	"context"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *AuthPresenter) AuthUserLogoutByToken(ctx context.Context, req *pb.AuthUserLogoutByTokenRequest) (*pb.AuthUserLogoutByTokenResponse, error) {
	ctx, span := p.telemetryInfrastructure.StartSpanFromContext(ctx, "AuthPresenter.AuthUserLogoutByToken")
	defer span.End()
	requestID := pkgContext.GetRequestIDFromContext(ctx)

	fullMethodName, err := pkgContext.GetFullMethodNameFromContext(ctx)
	if err != nil {
		p.logger.Error("UserPresenter.FindUserByEmailAndPassword", zap.String("requestID", requestID), zap.Error(err))
		return nil, err
	}

	// Access Control Authorization
	acl, err := p.authUseCase.AuthUserVerifyAccessControl(
		ctx,
		pkgContext.GetRequestIDFromContext(ctx),
		&pb.AuthUserVerifyAccessControlRequest{
			Token:          pkgContext.GetTokenAuthorizationFromContext(ctx),
			FullMethodName: &fullMethodName,
		},
	)
	if err != nil {
		p.logger.Error("UserPresenter.FindUserByEmailAndPassword", zap.String("requestID", requestID), zap.Error(err))
		return nil, nil
	}

	if !acl.Data.IsValid {
		p.logger.Error("UserPresenter.FindUserByEmailAndPassword", zap.String("requestID", requestID), zap.Error(err))
		return nil, nil
	}

	if err := req.Validate(); err != nil {
		p.logger.Error("AuthPresenter.AuthUserLogoutByToken", zap.String("requestID", requestID), zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err := p.authUseCase.AuthUserLogoutByToken(ctx, requestID, req)
	if err != nil {
		p.logger.Error("AuthPresenter.AuthUserLogoutByToken", zap.String("requestID", requestID), zap.Error(err))
		return nil, err
	}

	return resp, nil
}
