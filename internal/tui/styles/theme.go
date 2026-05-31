package styles

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	Title       lipgloss.Style
	ActiveTab   lipgloss.Style
	InactiveTab lipgloss.Style
	Content     lipgloss.Style
	Help        lipgloss.Style
}

func NewTheme() Theme {
	return Theme{
		Title:       lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("39")),
		ActiveTab:   lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("15")).Background(lipgloss.Color("62")).Padding(0, 1),
		InactiveTab: lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Padding(0, 1),
		Content:     lipgloss.NewStyle().Padding(1, 0),
		Help:        lipgloss.NewStyle().Foreground(lipgloss.Color("244")),
	}
}
