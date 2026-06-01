package app

import tea "github.com/charmbracelet/bubbletea"

func initialCmd(m AppModel) tea.Cmd {
	return tea.Batch(m.Docker.Init(), m.Systemd.Init(), m.Metrics.Init())
}
