package presenter

import (
	"context"

	pb "github.com/ferza17/ecommerce-microservices-v2/event-store-service/model/rpc/gen/v1/event"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *eventPresenter) ReadByAggregate(ctx context.Context, req *pb.ReadByAggregateRequest) (*pb.ReadByAggregateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadByAggregate not implemented")
}
