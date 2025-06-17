package pkg

import (
	"errors"
	"fmt"
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
		cb          *gobreaker.CircuitBreaker
		serviceName string
		logger      IZapLogger
	}
)

func (c *circuitBreaker) Execute(req func() (interface{}, error)) (interface{}, error) {
	result, err := c.cb.Execute(req)

	if err != nil {

		switch {
		case errors.Is(err, gobreaker.ErrOpenState):
			c.logger.Error(fmt.Sprintf("Circuit Breaker for %s is open. Request Failed: %v\n", c.serviceName, err))
			return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Circuit Breaker for %s is open. Please try again later", c.serviceName))
		case errors.Is(err, gobreaker.ErrTooManyRequests):
			c.logger.Error(fmt.Sprintf("Circuit Breaker for %s in half-open mode and too many request: %v\n", c.serviceName, err))
			return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Circuit Breaker for %s is open. Please try again later", c.serviceName))
		default:
			return nil, fmt.Errorf("failed to call %s: %w", c.serviceName, err)
		}
	}

	return result, nil
}

func NewCircuitBreaker(svcName string, logger IZapLogger) ICircuitBreaker {
	return &circuitBreaker{
		logger:      logger,
		serviceName: svcName,
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
