package systemdcomponent

import (
	"axiomtui/internal/domain"
	"axiomtui/internal/systemd"

	tea "github.com/charmbracelet/bubbletea"
)

// Model for the Systemd tab.
type Model struct {
	Services []domain.Service
	Error    error
	Loaded   bool
}

func NewModel() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return fetchServices()
}

func fetchServices() tea.Cmd {
	return func() tea.Msg {
		services, err := systemd.NewClient().ListServices()
		if err != nil {
			return ServicesError{Err: err}
		}

		return ServicesLoaded{Services: services}
	}
}
