package presenter

import (
	"context"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"go.uber.org/zap"
)

func (p *UserPresenter) FindUserByEmailAndPassword(ctx context.Context, req *userRpc.FindUserByEmailAndPasswordRequest) (*userRpc.User, error) {
	ctx, span := p.telemetryInfrastructure.Tracer(ctx, "UserPresenter.FindUserByEmailAndPassword")
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
		&userRpc.AuthUserVerifyAccessControlRequest{
			Token:          pkgContext.GetTokenAuthorizationFromContext(ctx),
			FullMethodName: &fullMethodName,
		},
	)
	if err != nil {
		p.logger.Error("UserPresenter.FindUserByEmailAndPassword", zap.String("requestID", requestID), zap.Error(err))
		return nil, nil
	}

	if !acl.IsValid {
		p.logger.Error("UserPresenter.FindUserByEmailAndPassword", zap.String("requestID", requestID), zap.Error(err))
		return nil, nil
	}

	res, err := p.userUseCase.FindUserByEmailAndPassword(ctx, requestID, req)
	if err != nil {
		p.logger.Error("UserPresenter.FindUserByEmailAndPassword", zap.String("requestID", requestID), zap.Error(err))
		return nil, err
	}
	return res, nil
}
