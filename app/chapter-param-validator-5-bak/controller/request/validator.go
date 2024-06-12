package request

import (
	"github.com/asaskevich/govalidator"
	"regexp"
)

func init() {
	govalidator.TagMap["IsMobile"] = func(value string) bool {
		return IsMobile(value)
	}
}

func IsMobile(value string) bool {
	matched, _ := regexp.MatchString(`^(1[3-9][0-9]\d{8})$`, value)
	return matched
}
