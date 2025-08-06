package temporal

import (
	"fmt"
	"github.com/nexus-rpc/sdk-go/nexus"
)

func (t *temporalInfrastructure) RegisterNexusOperation(operations ...nexus.RegisterableOperation) error {
	err := t.nexusService.Register(operations...)
	if err != nil {
		t.logger.Error(fmt.Sprintf("RegisterNexusOperation err: %v", err))
		return err
	}
	return nil
}
