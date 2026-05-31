package input

func IsNextTabKey(key string) bool {
	return key == "tab" || key == "right"
}

func IsPreviousTabKey(key string) bool {
	return key == "left"
}

func IsQuitKey(key string) bool {
	return key == "q" || key == "ctrl+c"
}
