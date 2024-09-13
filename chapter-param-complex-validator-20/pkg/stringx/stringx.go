package stringx

func Empty(s string) bool {
	return s == ""
}

func NotEmpty(s string) bool {
	return !Empty(s)
}
