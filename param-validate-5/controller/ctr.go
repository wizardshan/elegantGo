package controller

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"regexp"
)

var validate = binding.Validator.Engine().(*validator.Validate)

func init() {
	validate.RegisterValidation("mobile", func(fl validator.FieldLevel) bool {
		return isMobile(fl.Field().String())
	})
}

func isMobile(s string) bool {
	matched, _ := regexp.MatchString(`^(1[3-9][0-9]\d{8})$`, s)
	return matched
}
