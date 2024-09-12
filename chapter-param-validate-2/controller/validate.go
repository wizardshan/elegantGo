package controller

import "regexp"

func empty(s string) bool {
	return s == ""
}

func isMobile(s string) bool {
	matched, _ := regexp.MatchString(`^(1[3-9][0-9]\d{8})$`, s)
	return matched
}

func length(s string, value int) bool {
	return len(s) == value
}

func isNumber(s string) bool {
	matched, _ := regexp.MatchString(`^[0-9]+$`, s)
	return matched
}
