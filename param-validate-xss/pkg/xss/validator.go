package xss

import "errors"

var (
	ErrForbiddenKeywords = errors.New("forbidden keywords triggered, input contains one of the following keywords in a vulnerable format: alert, prompt")
)

var DefaultRules = []Rule{
	ForbiddenKeywords{},
	ForbiddenHTMLUnescapeStringKeywords{},
	ForbiddenURLQueryUnescapeKeywords{},
	ForbiddenUnicodeKeywords{},
	ForbiddenLowercaseKeywords{},
}

type Rule interface {
	Check(string) error
}

func Validate(input string) error {
	for _, rule := range DefaultRules {
		if err := rule.Check(input); err != nil {
			return err
		}
	}
	return nil
}

func ValidateRules(input string, rules ...Rule) error {
	for _, rule := range rules {
		if err := rule.Check(input); err != nil {
			return err
		}
	}
	return nil
}
