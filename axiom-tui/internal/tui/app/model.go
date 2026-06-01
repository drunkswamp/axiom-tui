package app

import (
	"os"

	"axiomtui/internal/localization"
	dockercomponent "axiomtui/internal/tui/components/docker"
	metricscomponent "axiomtui/internal/tui/components/metrics"
	systemdcomponent "axiomtui/internal/tui/components/systemd"
	"axiomtui/internal/tui/input"
	"axiomtui/internal/tui/styles"
)

type Tab int

const (
	DockerTab Tab = iota
	SystemdTab
	MetricsTab
)

const (
	dockerRegionID  = "docker"
	systemdRegionID = "systemd"
	metricsRegionID = "metrics"
)

type AppModel struct {
	ActiveTab  Tab
	Width      int
	Height     int
	Language   localization.Language
	Translator localization.Translator
	Theme      styles.Theme
	TabRegions []input.MouseRegion
	Docker     dockercomponent.Model
	Systemd    systemdcomponent.Model
	Metrics    metricscomponent.Model
}

// NewAppModel теперь возвращает (AppModel, error)
func NewAppModel() (AppModel, error) {
	language := localization.LanguageFromString(os.Getenv("AXIOMTUI_LANG"))

	docker, err := dockercomponent.NewModel()
	if err != nil {
		return AppModel{}, err
	}

	systemd := systemdcomponent.NewModel()

	metrics := metricscomponent.NewModel()

	return AppModel{
		ActiveTab:  DockerTab,
		Language:   language,
		Translator: localization.NewTranslator(language),
		Theme:      styles.NewTheme(),
		Docker:     docker,
		Systemd:    systemd,
		Metrics:    metrics,
	}, nil
}

func (m AppModel) NextTab() AppModel {
	m.ActiveTab = (m.ActiveTab + 1) % 3
	return m
}

func (m AppModel) PreviousTab() AppModel {
	if m.ActiveTab == DockerTab {
		m.ActiveTab = MetricsTab
		return m
	}

	m.ActiveTab--
	return m
}

func (m AppModel) WithTab(tab Tab) AppModel {
	m.ActiveTab = tab
	return m
}
