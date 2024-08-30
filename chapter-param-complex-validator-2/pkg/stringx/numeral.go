package stringx

import (
	"github.com/elliotchance/pie/v2"
	"regexp"
	"strconv"
)

// github.com/asaskevich/govalidator
func IsInt(s string) bool {
	if Empty(s) {
		return false
	}
	return regexp.MustCompile("^(?:[-+]?(?:0|[1-9][0-9]*))$").MatchString(s)
}

func ToInt(s string) int {
	return pie.Int(s)
}

func ToPtrInts(ss []string) []*int {
	return pie.Map(ss, func(s string) *int {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil
		}
		return &num
	})
}
