package controller

import (
	"elegantGo/chapter-param-validator-xss/pkg/xssvalidator"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
	"regexp"
)

var validate = binding.Validator.Engine().(*validator.Validate)

func init() {

	validate.RegisterValidation("mobile", func(fl validator.FieldLevel) bool {
		matched, _ := regexp.MatchString(`^(1[1-9][0-9]\d{8})$`, fl.Field().String())
		return matched
	})

	validate.RegisterValidation("xss", func(fl validator.FieldLevel) bool {
		err := xssvalidator.Validate(fl.Field().String())
		if err != nil {
			return false
		}

		return true
	})

	validate.RegisterValidation("nilfield", func(fl validator.FieldLevel) bool {
		field, _, _, ok := fl.GetStructFieldOK2()
		if !ok {
			return false
		}

		switch field.Kind() {
		case reflect.Slice, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func:
			return field.IsNil()
		default:
			return !field.IsValid() || field.IsZero()
		}
	})
}
