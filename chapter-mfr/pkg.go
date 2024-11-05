package main

import "regexp"

func IsInt(str string) bool {
	if str == "" {
		return false
	}
	return regexp.MustCompile("^(?:[-+]?(?:0|[1-9][0-9]*))$").MatchString(str)
}
