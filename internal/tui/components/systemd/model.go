package systemdcomponent

import "axiomtui/internal/domain"

// Model for the Systemd tab.
type Model struct {
	Services []domain.Service
	Error    error
}

func NewModel() Model {
	return Model{}
}
