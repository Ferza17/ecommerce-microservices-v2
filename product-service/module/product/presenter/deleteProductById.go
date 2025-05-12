package presenter

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
)

func (p *ProductGrpcPresenter) DeleteProductById(ctx context.Context, req *pb.DeleteProductByIdRequest) (*pb.DeleteProductByIdResponse, error) {
	ctx, span := p.telemetryInfrastructure.Tracer(ctx, "Presenter.DeleteProductById")
	defer span.End()
	return nil, nil
}
