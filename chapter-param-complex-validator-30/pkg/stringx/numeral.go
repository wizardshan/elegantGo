package stringx

import (
	"github.com/elliotchance/pie/v2"
	"regexp"
	"strconv"
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

func ToPtrInts(ss []string) []*int {
	return pie.Map(ss, func(s string) *int {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil
		}
		return &num
	})
}
