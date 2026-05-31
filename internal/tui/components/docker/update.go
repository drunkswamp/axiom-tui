package dockercomponent

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ServicesLoaded:
		m.Services = msg.Services
		m.Error = nil
	case ServicesError:
		m.Error = msg.Err
	}
	return m, nil
}
