// File: wire.go

//go:build wireinject
// +build wireinject

package presenter

import (
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
)

// ProvidePaymentPresenter wires dependencies for IPaymentPresenter.
func ProvidePaymentPresenter() IPaymentPresenter {
	wire.Build(
		NewPaymentPresenter,           // The constructor function for paymentPresenter
		usecase.ProvidePaymentUseCase, // Provides IPaymentUseCase
		telemetry.ProvideTelemetry,    // Provides ITelemetryInfrastructure
		logger.ProvideLogger,          // Provides IZapLogger
	)
	return nil // Wire will generate the concrete implementation.
}
