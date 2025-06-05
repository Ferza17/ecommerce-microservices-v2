// File: wire.go

//go:build wireinject
// +build wireinject

package presenter

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/module/provider/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
)

// PaymentProviderPresenterSet is a provider set for the NewPaymentProviderPresenter and its dependencies.
var PaymentProviderPresenterSet = wire.NewSet(
	usecase.ProvidePaymentProviderUseCase, // Provides IPaymentProviderUseCase
	telemetry.ProvideTelemetry,            // Provides ITelemetryInfrastructure
	logger.ProvideLogger,                  // Provides IZapLogger
	NewPaymentProviderPresenter,           // Provides IPaymentProviderPresenter
)

// ProvidePaymentProviderPresenter initializes and returns an IPaymentProviderPresenter.
func ProvidePaymentProviderPresenter(ctx context.Context) (IPaymentProviderPresenter, error) {
	wire.Build(PaymentProviderPresenterSet)
	return nil, nil
}
