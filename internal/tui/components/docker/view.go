package dockercomponent

import (
	"axiomtui/internal/localization"
	"fmt"
	"strings"
)

func (m Model) View(translator localization.Translator) string {
	if m.Error != nil {
		return fmt.Sprintf("Error: %v", m.Error)
	}

	if len(m.Services) == 0 {
		return translator.T(localization.KeyDockerPlaceholder)
	}

	var b strings.Builder
	b.WriteString("ID\t\tNAME\t\tSTATE\t\tIMAGE\n")
	for _, service := range m.Services {
		b.WriteString(fmt.Sprintf("%s\t%s\t\t%s\t\t%s\n",
			service.ID,
			service.Name,
			service.State,
			service.Type, // Type field from domain.Service holds the image name
		))
	}
	return b.String()
}
