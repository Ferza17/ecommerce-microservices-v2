//go:build wireinject
// +build wireinject

package circuit

import (
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
)

// InitializeCircuitBreaker is the Wire injector function
// that provides an instance of ICircuitBreaker.
func InitializeCircuitBreaker(svcName string, logger logger.IZapLogger) ICircuitBreaker {
	wire.Build(NewCircuitBreaker)
	return nil
}
