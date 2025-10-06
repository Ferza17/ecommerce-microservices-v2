package presenter

import (
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	userUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/usecase"
	"github.com/google/wire"

	"github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
)

type AuthPresenter struct {
	userRpc.UnimplementedAuthServiceServer

	authUseCase             usecase.IAuthUseCase
	userUseCase             userUseCase.IUserUseCase
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
	logger                  logger.IZapLogger
}

var Set = wire.NewSet(NewAuthPresenter)

func NewAuthPresenter(
	authUseCase usecase.IAuthUseCase,
	userUseCase userUseCase.IUserUseCase,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger) *AuthPresenter {
	c := &AuthPresenter{
		authUseCase:             authUseCase,
		userUseCase:             userUseCase,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
	return c
}
