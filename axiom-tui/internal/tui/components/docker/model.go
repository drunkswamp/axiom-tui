package dockercomponent

import (
	"axiomtui/internal/domain"
	"axiomtui/internal/docker" // Import the docker adapter
	tea "github.com/charmbracelet/bubbletea"
	"context"
)

// Model for the Docker tab now holds a list of unified services.
type Model struct {
	Services []domain.Service
	Error    error
	adapter  *docker.Adapter // Add docker adapter
}

func NewModel() (Model, error) {
	adapter, err := docker.NewAdapter()
	if err != nil {
		return Model{}, err
	}

	return Model{
		adapter: adapter,
	}, nil
}

// Init is the first command that will be run when the model is focused.
func (m Model) Init() tea.Cmd {
	return fetchServices(m.adapter)
}

// A command to fetch the services.
func fetchServices(adapter *docker.Adapter) tea.Cmd {
	return func() tea.Msg {
		services, err := adapter.GetServices(context.Background())
		if err != nil {
			return ServicesError{Err: err}
		}
		return ServicesLoaded{Services: services}
	}
}
