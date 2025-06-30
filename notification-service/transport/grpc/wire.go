//go:build wireinject
// +build wireinject

package grpc

import (
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	"github.com/google/wire"
)

func ProvideGrpcServer() *GrpcServer {
	wire.Build(
		logger.Set,
		Set)
	return nil
}
