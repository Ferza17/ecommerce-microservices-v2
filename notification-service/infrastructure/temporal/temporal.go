package temporal

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	"github.com/google/wire"
	"github.com/nexus-rpc/sdk-go/nexus"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

type (
	ITemporalInfrastructure interface {
		StartWorkflow(ctx context.Context, workflowID string, workflowType interface{}, args ...interface{}) (client.WorkflowRun, error)
		SignalWorkflow(ctx context.Context, workflowID string, signalName string, signalValue interface{}) error
		//GetWorkflowResult(ctx context.Context, workflowID string, runID string, valuePtr interface{}) error
		Start() error
		//Stop()

		// Nexus
		RegisterNexusOperation(operations ...nexus.RegisterableOperation) error

		RegisterWorkflow(w interface{}) ITemporalInfrastructure
		RegisterActivity(a interface{}) ITemporalInfrastructure

		// Register Nexus Operation
	}

	temporalInfrastructure struct {
		logger       logger.IZapLogger
		client       client.Client
		worker       worker.Worker
		nexusService *nexus.Service
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
		client: c,
		worker: worker.New(c, config.Get().NotificationServiceServiceName, worker.Options{
			EnableLoggingInReplay:       true,
			Identity:                    config.Get().NotificationServiceServiceName,
			DisableRegistrationAliasing: true,
		}),
		nexusService: nexus.NewService(config.Get().NotificationServiceServiceName),
	}
}
