package controller

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	NotEmpty = "NotEmpty"
	IsMobile = "IsMobile"
	Length   = "Length"
	IsNumber = "IsNumber"
)

type validateErrs []error

func (errs validateErrs) Error() string {
	errsMsg := ""
	for _, err := range errs {
		errsMsg += err.Error()
	}
	return errsMsg
}

type field struct {
	name      string
	funcNames []string
	value     string
	args      []any
}

func validate(fields []*field) (errs validateErrs) {
	for _, f := range fields {
		if es := valid(f); es != nil {
			errs = append(errs, es...)
		}
	}
	return
}

func valid(field *field) (errs []error) {
	for _, funcName := range field.funcNames {
		errMsg := ""
		switch funcName {
		case NotEmpty:
			if empty(field.value) {
				errMsg = field.name + "不能为空"
			}
		case IsMobile:
			if !isMobile(field.value) {
				errMsg = "手机号格式不正确"
			}
		case Length:
			v := field.args[0].(int)
			if !length(field.value, v) {
				errMsg = fmt.Sprintf("%s必须%d位", field.name, v)
			}
		case IsNumber:
			if !isNumber(field.value) {
				errMsg = field.name + "必须数字"
			}
		default:
			errMsg = funcName + "校验函数不存在"
		}

		if errMsg == "" {
			continue
		}
		errs = append(errs, errors.New(errMsg))
	}
	return
}

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
