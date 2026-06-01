package metricscomponent

import (
	"axiomtui/internal/metrics"

	tea "github.com/charmbracelet/bubbletea"
)

// Model for the Metrics tab.
type Model struct {
	Snapshot metrics.Snapshot
	Loaded   bool
	Error    error
}

func NewModel() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return fetchSnapshot()
}

func fetchSnapshot() tea.Cmd {
	return func() tea.Msg {
		snapshot, err := metrics.NewCollector().Snapshot()
		if err != nil {
			return SnapshotError{Err: err}
		}

		return SnapshotLoaded{Snapshot: snapshot}
	}
}
