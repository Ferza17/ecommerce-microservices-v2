package telemetry

import (
	"context"
	"fmt"
)

func (t *telemetryInfrastructure) Shutdown(ctx context.Context) error {
	if err := t.tracerProvider.Shutdown(ctx); err != nil {
		t.logger.Error(fmt.Sprintf("Failed to shutdown tracer provider: %v", err))
		return err
	}
	return nil
}
