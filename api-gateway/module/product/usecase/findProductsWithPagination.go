package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/enum"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/pb"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/util"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc/metadata"
)

func (u *ProductUseCase) FindProductsWithPagination(ctx context.Context, requestId string, req *pb.FindProductsWithPaginationRequest) (*pb.FindProductsWithPaginationResponse, error) {
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.FindProductsWithPagination")
	defer span.End()

	md := metadata.New(map[string]string{enum.XRequestIDHeader.String(): requestId})
	otel.GetTextMapPropagator().Inject(ctx, &util.MetadataHeaderCarrier{md})
	ctx = metadata.NewOutgoingContext(ctx, md)

	products, err := u.rpcClient.GetProductService().FindProductsWithPagination(ctx, req)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error finding products with pagination: %v", err))
		return nil, err
	}

	return products, nil
}
