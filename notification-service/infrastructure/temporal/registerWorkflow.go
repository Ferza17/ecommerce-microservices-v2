package temporal

func (t *temporalInfrastructure) RegisterWorkflow(w interface{}) ITemporalInfrastructure {
	t.worker.RegisterWorkflow(w)
	return t
}
