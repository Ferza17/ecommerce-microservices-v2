package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/enum"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/service"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/pb"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc/metadata"
)

func (u *ProductUseCase) FindProductById(ctx context.Context, requestId string, req *pb.FindProductByIdRequest) (*pb.Product, error) {
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.FindProductById")
	defer span.End()

	md := metadata.New(map[string]string{enum.XRequestIDHeader.String(): requestId})
	otel.GetTextMapPropagator().Inject(ctx, &service.MetadataHeaderCarrier{md})
	ctx = metadata.NewOutgoingContext(ctx, md)
	product, err := u.rpcClient.GetProductService().FindProductById(ctx, req)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error finding product by id: %v", err))
		return nil, err
	}

	return product, nil
}
