//go:build wireinject
// +build wireinject

package telemetry

import (
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
)

// ProvideTelemetry wires dependencies for ITelemetryInfrastructure.
func ProvideTelemetry(logger logger.IZapLogger) ITelemetryInfrastructure {
	wire.Build(NewTelemetry)
	return NewTelemetry(logger) // Wire will generate the actual implementation.
}
