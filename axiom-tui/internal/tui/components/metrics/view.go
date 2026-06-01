package metricscomponent

import (
	"axiomtui/internal/localization"
	"fmt"
)

func (m Model) View(translator localization.Translator) string {
	if m.Error != nil {
		return fmt.Sprintf(translator.T(localization.KeyMetricsErrorFormat), m.Error)
	}

	if !m.Loaded {
		return translator.T(localization.KeyMetricsLoading)
	}

	return fmt.Sprintf(
		translator.T(localization.KeyMetricsSummary),
		m.Snapshot.CPUUsage,
		m.Snapshot.MemoryUsage,
	)
}
