package dockercomponent

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ServicesLoaded:
		m.Services = msg.Services
		m.Error = nil
		m.Loaded = true
	case ServicesError:
		m.Error = msg.Err
		m.Loaded = true
	}
	return m, nil
}
