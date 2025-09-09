package circuit

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type (
	ICircuitBreaker interface {
		Execute(req func() (interface{}, error)) (interface{}, error)
	}

	circuitBreaker struct {
		cb     *gobreaker.CircuitBreaker
		logger logger.IZapLogger
	}
)

// Set is a Wire provider set for circuit breaker dependencies
var Set = wire.NewSet(
	NewCircuitBreaker,
	wire.Bind(new(ICircuitBreaker), new(*circuitBreaker)),
	ProvideServiceName,
)

// ProvideServiceName provides the service name for circuit breaker
func ProvideServiceName() string {
	return config.Get().PaymentServiceServiceName
}

func NewCircuitBreaker(svcName string, logger logger.IZapLogger) ICircuitBreaker {
	return &circuitBreaker{
		logger: logger,
		cb: gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:        svcName,
			MaxRequests: 3,
			Interval:    5 * time.Second,
			Timeout:     3 * time.Second,
			ReadyToTrip: func(counts gobreaker.Counts) bool {
				failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
				return counts.Requests >= 3 && failureRatio >= 0.6
			},
			OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
				fmt.Printf("Circuit Breaker '%s' changed state from %s to %s\n", name, from, to)
			},
		}),
	}
}

func (c *circuitBreaker) Execute(req func() (interface{}, error)) (interface{}, error) {
	result, err := c.cb.Execute(req)

	if err != nil {
		if err == gobreaker.ErrOpenState {
			c.logger.Error(fmt.Sprintf("Circuit Breaker for User Service is open. Request Failed: %v\n", err))
			return nil, status.Errorf(codes.Unavailable, "User Service is currently unavailable")
		}
		if err == gobreaker.ErrTooManyRequests {
			c.logger.Error(fmt.Sprintf("Circuit Breaker for User Service in half-open mode and too many request: %v\n", err))
			return nil, status.Errorf(codes.Unavailable, "User Service is busy, please try again later")
		}
		return nil, fmt.Errorf("failed to call User Service: %w", err)
	}

	return result, nil
}
