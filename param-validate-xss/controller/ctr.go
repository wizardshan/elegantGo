package controller

import (
	"elegantGo/param-validate-xss/pkg/xss"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"regexp"
)

var validate = binding.Validator.Engine().(*validator.Validate)

func init() {
	validate.RegisterValidation("sqlinject", func(fl validator.FieldLevel) bool {
		return !sqlInject(fl.Field().String())
	})

	validate.RegisterValidation("xss", func(fl validator.FieldLevel) bool {
		err := xss.Validate(fl.Field().String())
		if err != nil {
			return false
		}
		return true
	})
}

func sqlInject(s string) bool {
	detectSqlInjectRegex := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
	return regexp.MustCompile(detectSqlInjectRegex).MatchString(s)
}
