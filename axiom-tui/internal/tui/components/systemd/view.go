package systemdcomponent

import (
	"axiomtui/internal/localization"
	"fmt"
	"strings"
)

func (m Model) View(translator localization.Translator) string {
	if m.Error != nil {
		return fmt.Sprintf(translator.T(localization.KeySystemdErrorFormat), m.Error)
	}

	if !m.Loaded {
		return translator.T(localization.KeySystemdLoading)
	}

	if len(m.Services) == 0 {
		return translator.T(localization.KeySystemdEmpty)
	}

	var builder strings.Builder
	builder.WriteString(translator.T(localization.KeySystemdTableHeader))
	for _, service := range m.Services {
		fmt.Fprintf(&builder, "\n%s\t%s\t%s", service.ID, service.Name, service.State.Status)
	}

	return builder.String()
}
