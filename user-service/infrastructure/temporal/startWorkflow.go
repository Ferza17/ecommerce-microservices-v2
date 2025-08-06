package temporal

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"go.temporal.io/sdk/client"
)

func (t *temporalInfrastructure) StartWorkflow(ctx context.Context, workflowID string, workflowType interface{}, args ...interface{}) (client.WorkflowRun, error) {
	options := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: config.Get().UserServiceServiceName,
	}
	return t.client.ExecuteWorkflow(ctx, options, workflowType, args...)
}
