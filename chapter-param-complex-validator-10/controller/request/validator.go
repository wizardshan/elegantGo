package request

import (
	"elegantGo/chapter-param-complex-validator-1/pkg/stringx"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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
}
