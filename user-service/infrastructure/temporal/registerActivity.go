package temporal

func (t *temporalInfrastructure) RegisterActivity(a interface{}) ITemporalInfrastructure {
	t.worker.RegisterActivity(a)
	return t
}
