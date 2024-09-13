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

var (
	ErrNotEmptyMsg  = "%s不能为空"
	ErrNotNumberMsg = "%s必须数字"
	ErrFormatMsg    = "%s格式不正确"
	ErrLengthMsg    = "%s必须%d位"

	ErrValidateFuncNotExistMsg = "%s校验函数不存在"
)

type Validator func() bool

var ValidatorFuncMapping = map[string]func(field *Field) (Validator, string){
	NotEmpty: func(field *Field) (Validator, string) {
		return func() bool {
			return notEmpty(field.value)
		}, fmt.Sprintf(ErrNotEmptyMsg, field.name)
	},
	IsMobile: func(field *Field) (Validator, string) {
		return func() bool {
			return isMobile(field.value)
		}, fmt.Sprintf(ErrFormatMsg, field.name)
	},
	Length: func(field *Field) (Validator, string) {
		v := field.args[0].(int)
		return func() bool {
			return length(field.value, v)
		}, fmt.Sprintf(ErrLengthMsg, field.name, v)
	},
	IsNumber: func(field *Field) (Validator, string) {
		return func() bool {
			return isNumber(field.value)
		}, fmt.Sprintf(ErrNotNumberMsg, field.name)
	},
}

type Errs []error

func (errs Errs) Error() string {
	errsMsg := ""
	for _, err := range errs {
		errsMsg += err.Error()
	}
	return errsMsg
}

type Fields []*Field

func (fields Fields) validate() (errs Errs) {
	for _, field := range fields {
		if err := field.validate(); err != nil {
			errs = append(errs, err)
		}
	}
	return
}

type Field struct {
	name      string
	funcNames []string
	value     string
	args      []any
}

func (field *Field) validate() (errs Errs) {
	for _, funcName := range field.funcNames {
		vf, ok := ValidatorFuncMapping[funcName]
		if !ok {
			errs = append(errs, fmt.Errorf(ErrValidateFuncNotExistMsg, field.name))
			continue
		}

		validator, errMsg := vf(field)
		// 前文中所有的if控制语句都抽取成这一个if代码块了，并且重复使用也遗漏不了!符号
		if !validator() {
			errs = append(errs, errors.New(errMsg))
		}
	}
	return
}

func notEmpty(s string) bool {
	return s != ""
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
