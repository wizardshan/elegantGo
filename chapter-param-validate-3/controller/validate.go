package controller

import (
	"errors"
	"regexp"
)

var (
	ErrMobileEmpty     = errors.New("手机号不能为空")
	ErrMobileNotNumber = errors.New("手机号必须数字")
	ErrMobileFormat    = errors.New("手机号格式不正确")

	ErrCaptchaEmpty     = errors.New("验证码不能为空")
	ErrCaptchaNotNumber = errors.New("验证码必须数字")
	ErrCaptchaLength    = errors.New("验证码必须4位")
)

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
