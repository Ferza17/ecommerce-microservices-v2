package pkg

import (
	"fmt"
	"github.com/sony/gobreaker"
	"time"
)

type (
	ICircuitBreaker interface {
		Execute(func() (interface{}, error)) (interface{}, error)
	}

	circuitBreaker struct {
		cb *gobreaker.CircuitBreaker
	}
)

func (c *circuitBreaker) Execute(f func() (interface{}, error)) (interface{}, error) {
	return c.Execute(f)
}

func NewCircuitBreaker(svcName string) ICircuitBreaker {
	return &circuitBreaker{
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

func (c *circuitBreaker) GetCB() *gobreaker.CircuitBreaker {
	return c.cb
}
