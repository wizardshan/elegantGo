package controller

import (
	"fmt"
	"regexp"
)

const (
	NotEmpty = "NotEmpty"
	IsMobile = "IsMobile"
	Length   = "Length"
	IsNumber = "IsNumber"
)

var (
	ErrEmptyMsg             = "%s不能为空"
	ErrNotNumberMsg         = "%s必须数字"
	ErrFormatMsg            = "%s格式不正确"
	ErrLengthMsg            = "%s必须%d位"
	ErrValidateFuncNotExist = "%s校验函数不存在"
)

type Errs []error

func (errs Errs) Error() string {
	errsMsg := ""
	for _, err := range errs {
		errsMsg += err.Error()
	}
	return errsMsg
}

type Field struct {
	name  string
	funcs []string
	value string
	args  []any
}

func validate(fields []*Field) (errs Errs) {
	for _, field := range fields {
		for _, f := range field.funcs {
			var err error
			switch f {
			case NotEmpty:
				if empty(field.value) {
					err = fmt.Errorf(ErrEmptyMsg, field.name)
				}
			case IsMobile:
				if !isMobile(field.value) {
					err = fmt.Errorf(ErrFormatMsg, field.name)
				}
			case Length:
				v := field.args[0].(int)
				if !length(field.value, v) {
					err = fmt.Errorf(ErrLengthMsg, field.name, v)
				}
			case IsNumber:
				if !isNumber(field.value) {
					err = fmt.Errorf(ErrNotNumberMsg, field.name)
				}
			default:
				err = fmt.Errorf(ErrValidateFuncNotExist, field.name)
			}

			if err == nil {
				continue
			}
			errs = append(errs, err)
		}
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
