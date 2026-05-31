package app

import (
	"strings"

	"axiomtui/internal/localization"
	"axiomtui/internal/tui/input"
	"axiomtui/internal/tui/styles"
)

func (m AppModel) View() string {
	title := m.Theme.Title.Render(m.Translator.T(localization.KeyAppTitle))
	tabs, regions := m.renderTabs()
	m.TabRegions = regions
	content := m.Theme.Content.Render(m.activeContent())
	help := styles.RenderHelp(m.Theme, m.Translator.T(localization.KeyHelp))

	return strings.Join([]string{title, tabs, content, help}, "\n")
}

func (m AppModel) renderTabs() (string, []input.MouseRegion) {
	items := []struct {
		id    string
		tab   Tab
		label string
	}{
		{id: dockerRegionID, tab: DockerTab, label: m.Translator.T(localization.KeyTabDocker)},
		{id: systemdRegionID, tab: SystemdTab, label: m.Translator.T(localization.KeyTabSystemd)},
		{id: metricsRegionID, tab: MetricsTab, label: m.Translator.T(localization.KeyTabMetrics)},
	}

	var rendered []string
	var regions []input.MouseRegion
	x := 0

	for _, item := range items {
		style := m.Theme.InactiveTab
		if item.tab == m.ActiveTab {
			style = m.Theme.ActiveTab
		}

		renderedTab := style.Render(item.label)
		rendered = append(rendered, renderedTab)
		regions = append(regions, input.MouseRegion{
			ID:     item.id,
			X:      x,
			Y:      1,
			Width:  len(item.label) + 2,
			Height: 1,
		})
		x += len(item.label) + 3
	}

	return strings.Join(rendered, " "), regions
}

func (m AppModel) activeContent() string {
	switch m.ActiveTab {
	case DockerTab:
		return m.Docker.View(m.Translator)
	case SystemdTab:
		return m.Systemd.View(m.Translator)
	case MetricsTab:
		return m.Metrics.View(m.Translator)
	default:
		return ""
	}
}
