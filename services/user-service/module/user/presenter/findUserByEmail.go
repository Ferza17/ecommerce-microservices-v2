package presenter

import (
	"context"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"go.uber.org/zap"
)

func (p *UserPresenter) FindUserByEmail(ctx context.Context, req *pb.FindUserByEmailRequest) (*pb.FindUserByEmailResponse, error) {
	ctx, span := p.telemetryInfrastructure.StartSpanFromContext(ctx, "UserPresenter.FindUserByEmail")
	defer span.End()
	requestID := pkgContext.GetRequestIDFromContext(ctx)

	res, err := p.userUseCase.FindUserByEmail(ctx, requestID, req)
	if err != nil {
		p.logger.Error("UserPresenter.FindUserByEmailAndPassword", zap.String("requestID", requestID), zap.Error(err))
		return nil, err
	}
	return res, nil
}
