package presenter

import (
	"context"

	pb "github.com/ferza17/ecommerce-microservices-v2/event-store-service/model/rpc/gen/v1/event"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/context"
	"go.uber.org/zap"
)

func (p *eventPresenter) Append(ctx context.Context, req *pb.AppendRequest) (*pb.AppendResponse, error) {
	ctx, span := p.telemetry.StartSpanFromContext(ctx, "EventPresenter.Append")
	defer span.End()

	if err := req.Validate(); err != nil {
		p.logger.Error("EventPresenter.Append", zap.String("requestID", pkgContext.GetRequestIDFromContext(ctx)), zap.Error(err))
		return nil, err
	}

	resp, err := p.eventUseCase.Append(ctx, pkgContext.GetRequestIDFromContext(ctx), req)
	if err != nil {
		p.logger.Error("EventPresenter.Append", zap.Error(err))
		return nil, err
	}

	return resp, nil
}
