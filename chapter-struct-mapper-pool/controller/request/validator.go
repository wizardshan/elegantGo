package request

import (
	xssvalidator "app/chapter-struct-mapper-pool/pkg/xss-validator"
	"github.com/asaskevich/govalidator"
	"regexp"
)

func init() {
	govalidator.TagMap["IsMobile"] = func(value string) bool {
		return IsMobile(value)
	}
	govalidator.TagMap["CheckXSS"] = func(value string) bool {
		return CheckXSS(value)
	}
}

func IsMobile(value string) bool {
	matched, _ := regexp.MatchString(`^(1[3-9][0-9]\d{8})$`, value)
	return matched
}

func CheckXSS(value string) bool {
	err := xssvalidator.Validate(value)
	if err != nil {
		return false
	}

	return true
}
