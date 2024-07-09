package controller

import (
	"elegantGo/chapter-param-validator-xss/pkg/xssvalidator"
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

	validate.RegisterValidation("xss", func(fl validator.FieldLevel) bool {
		err := xssvalidator.Validate(fl.Field().String())
		if err != nil {
			return false
		}

		return true
	})
}
