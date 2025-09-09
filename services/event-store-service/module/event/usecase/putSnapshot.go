package usecase

import (
	"context"

	pb "github.com/ferza17/ecommerce-microservices-v2/event-store-service/model/rpc/gen/v1/event"
)

func (u *eventUseCase) PutSnapshot(ctx context.Context, requestId string, req *pb.PutSnapshotRequest) (*pb.PutSnapshotResponse, error) {
	//TODO implement me
	panic("implement me")
}
