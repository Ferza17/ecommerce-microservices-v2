package presenter

import (
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/temporal"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"github.com/google/wire"

	"github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
)

type AuthPresenter struct {
	userRpc.UnimplementedAuthServiceServer

	authUseCase             usecase.IAuthUseCase
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
	temporal                temporal.ITemporalInfrastructure
	logger                  logger.IZapLogger
}

var Set = wire.NewSet(NewAuthPresenter)

func NewAuthPresenter(
	authUseCase usecase.IAuthUseCase,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	temporal temporal.ITemporalInfrastructure,
	logger logger.IZapLogger) *AuthPresenter {
	return &AuthPresenter{
		authUseCase:             authUseCase,
		telemetryInfrastructure: telemetryInfrastructure,
		temporal:                temporal,
		logger:                  logger,
	}
}
