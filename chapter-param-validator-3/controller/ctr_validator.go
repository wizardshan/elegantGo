package controller

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"log"
	"reflect"
	"regexp"
)

var validate = binding.Validator.Engine().(*validator.Validate)
var trans ut.Translator

func init() {
	validate.RegisterValidation("mobile", func(fl validator.FieldLevel) bool {
		matched, _ := regexp.MatchString(`^(1[1-9][0-9]\d{8})$`, fl.Field().String())
		return matched
	})

	zh := zh.New()
	var found bool
	trans, found = ut.New(zh, zh).GetTranslator(zh.Locale())
	if !found {
		log.Fatal("translator not found")
	}
	zhTranslations.RegisterDefaultTranslations(validate, trans)

	// 设置错误提示语
	validate.RegisterTranslation("mobile", trans, func(ut ut.Translator) error {
		return ut.Add("mobile", "{0}格式不正确", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("mobile", fe.Field())
		return t
	})

	// 设置字段名名称
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		tagName := field.Name
		label := field.Tag.Get("label")
		if label != "" {
			tagName = label
		}
		return tagName
	})
}
