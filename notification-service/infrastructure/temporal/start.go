package temporal

func (t *temporalInfrastructure) Start() error {
	t.logger.Info("Starting Temporal worker...")
	t.worker.RegisterNexusService(t.nexusService)
	return t.worker.Start()
}
