package temporal

import "context"

func (t *temporalInfrastructure) SignalWorkflow(ctx context.Context, workflowID string, signalName string, signalValue interface{}) error {
	return t.client.SignalWorkflow(ctx, workflowID, "", signalName, signalValue)
}
