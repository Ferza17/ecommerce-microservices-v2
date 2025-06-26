package presenter

import (
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"github.com/google/wire"

	"github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
)

type UserPresenter struct {
	userRpc.UnimplementedUserServiceServer

	userUseCase             usecase.IUserUseCase
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
	logger                  logger.IZapLogger
}

var Set = wire.NewSet(NewUserPresenter)

func NewUserPresenter(
	userUseCase usecase.IUserUseCase,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger) *UserPresenter {
	return &UserPresenter{
		userUseCase:             userUseCase,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
