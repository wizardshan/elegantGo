package xssvalidator

import "errors"

var (
	ErrForbiddenKeywords = errors.New("forbidden keywords triggered, input contains one of the following keywords in a vulnerable format: alert, prompt")
)
