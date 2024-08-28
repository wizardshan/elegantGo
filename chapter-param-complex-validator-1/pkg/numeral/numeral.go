package numeral

import (
	"github.com/elliotchance/pie/v2"
	"regexp"
)

// github.com/asaskevich/govalidator
func IsInt(str string) bool {
	if str == "" {
		return false
	}
	return regexp.MustCompile("^(?:[-+]?(?:0|[1-9][0-9]*))$").MatchString(str)
}

func ToInt(s string) int {
	return pie.Int(s)
}
