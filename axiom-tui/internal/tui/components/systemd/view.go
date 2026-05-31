package systemdcomponent

import (
	"axiomtui/internal/localization"
	"fmt"
)

func (m Model) View(translator localization.Translator) string {
	if m.Error != nil {
		return fmt.Sprintf("Ошибка: %v", m.Error)
	}

	if len(m.Services) == 0 {
		return translator.T(localization.KeySystemdPlaceholder)
	}

	return fmt.Sprintf("Найдено сервисов: %d", len(m.Services))
}
