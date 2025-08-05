package temporal

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

type (
	ITemporalInfrastructure interface {
	}

	temporalInfrastructure struct {
		logger logger.IZapLogger
		worker worker.Worker
	}
)

var Set = wire.NewSet(NewTemporalInfrastructure)

func NewTemporalInfrastructure(logger logger.IZapLogger) ITemporalInfrastructure {
	c, err := client.Dial(client.Options{
		HostPort: "localhost:7233",
	})
	if err != nil {
		logger.Error(fmt.Sprintf("failed to connect to temporal: %v", err))
		panic(err)
	}
	
	return &temporalInfrastructure{
		logger: logger,
		worker: worker.New(c, config.Get().UserServiceServiceName, worker.Options{}),
	}
}
