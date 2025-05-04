package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/pb"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/gatewayEventStore/repository/mongodb"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
)

type (
	IGatewayEventStoreUseCase interface {
		CreateGatewayEventStore(ctx context.Context, requestId string, req *pb.CreateGatewayEventStoreRequest) (*pb.CreateGatewayEventStoreResponse, error)
		CreateGatewayEventStoreWithTransaction(ctx context.Context, requestId string, req *pb.CreateGatewayEventStoreRequest) (*pb.CreateGatewayEventStoreResponse, error)
	}

	gatewayEventStoreUseCase struct {
		gatewayEventStoreRepository mongodb.IGatewayEventStoreRepository
		logger                      pkg.IZapLogger
	}
)

func NewGatewayEventStoreUseCase(gatewayEventStoreRepository mongodb.IGatewayEventStoreRepository, logger pkg.IZapLogger) IGatewayEventStoreUseCase {
	return &gatewayEventStoreUseCase{
		gatewayEventStoreRepository: gatewayEventStoreRepository,
		logger:                      logger,
	}
}
