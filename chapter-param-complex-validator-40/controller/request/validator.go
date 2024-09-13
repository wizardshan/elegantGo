package request

import (
	"elegantGo/chapter-param-complex-validator-4/pkg/stringx"
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

	validate.RegisterValidation("int", func(fl validator.FieldLevel) bool {
		return stringx.IsInt(fl.Field().String())
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
