package dockercomponent

import (
	"axiomtui/internal/localization"
	"fmt"
	"strings"
)

func (m Model) View(translator localization.Translator) string {
	if m.Error != nil {
		return fmt.Sprintf(translator.T(localization.KeyDockerErrorFormat), m.Error)
	}

	if !m.Loaded {
		return translator.T(localization.KeyDockerLoading)
	}

	if len(m.Services) == 0 {
		return translator.T(localization.KeyDockerEmpty)
	}

	var b strings.Builder
	b.WriteString(translator.T(localization.KeyDockerTableHeader))
	for _, service := range m.Services {
		fmt.Fprintf(&b, "%s\t%s\t\t%s\t\t%s\n",
			service.ID,
			service.Name,
			service.State.Status,
			service.Type,
		)
	}
	return b.String()
}
