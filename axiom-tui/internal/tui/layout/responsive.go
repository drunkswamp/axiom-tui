package layout

func ContentWidth(windowWidth int) int {
	if windowWidth < 0 {
		return 0
	}

	return windowWidth
}
