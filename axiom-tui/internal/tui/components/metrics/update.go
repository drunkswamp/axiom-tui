package metricscomponent

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case SnapshotLoaded:
		m.Snapshot = msg.Snapshot
		m.Loaded = true
		m.Error = nil
	case SnapshotError:
		m.Loaded = true
		m.Error = msg.Err
	}

	return m, nil
}
