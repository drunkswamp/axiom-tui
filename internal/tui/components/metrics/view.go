package metricscomponent

import (
	"axiomtui/internal/localization"
	"fmt"
)

func (m Model) View(translator localization.Translator) string {
	if m.Error != nil {
		return fmt.Sprintf("Error: %v", m.Error)
	}

	if m.Snapshot.Timestamp.IsZero() {
		return translator.T(localization.KeyMetricsPlaceholder)
	}

	return fmt.Sprintf(
		"CPU: %.2f%% | RAM: %.2f%%",
		m.Snapshot.CPUUsage,
		m.Snapshot.MemoryUsage,
	)
}
