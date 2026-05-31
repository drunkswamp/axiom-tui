package app

import (
	"axiomtui/internal/tui/input"

	tea "github.com/charmbracelet/bubbletea"
)

func (m AppModel) Init() tea.Cmd {
	return initialCmd()
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		return m, nil
	case tea.KeyMsg:
		key := msg.String()
		if input.IsQuitKey(key) {
			return m, tea.Quit
		}
		if input.IsNextTabKey(key) {
			return m.NextTab(), nil
		}
		if input.IsPreviousTabKey(key) {
			return m.PreviousTab(), nil
		}
	case tea.MouseMsg:
		if msg.Type != tea.MouseLeft {
			return m, nil
		}
		for _, region := range m.TabRegions {
			if !region.Contains(msg.X, msg.Y) {
				continue
			}
			switch region.ID {
			case dockerRegionID:
				return m.WithTab(DockerTab), nil
			case systemdRegionID:
				return m.WithTab(SystemdTab), nil
			case metricsRegionID:
				return m.WithTab(MetricsTab), nil
			}
		}
	}

	return m, nil
}
