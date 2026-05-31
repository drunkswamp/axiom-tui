package styles

func RenderHelp(theme Theme, help string) string {
	return theme.Help.Render(help)
}
