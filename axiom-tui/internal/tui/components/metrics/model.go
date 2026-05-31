package metricscomponent

import "axiomtui/internal/metrics"

// Model for the Metrics tab.
type Model struct {
	Snapshot metrics.Snapshot
	Error    error
}

func NewModel() Model {
	return Model{}
}
