package usecase

import (
	"context"

	pb "github.com/ferza17/ecommerce-microservices-v2/event-store-service/model/rpc/gen/v1/event"
)

func (u *eventUseCase) Append(ctx context.Context, requestId string, req *pb.AppendRequest) (*pb.AppendResponse, error) {
	//TODO implement me
	panic("implement me")
}
