package gin

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

var (
	trans    ut.Translator
	Validate = binding.Validator.Engine().(*validator.Validate)
)

func init() {

	// 自定义手机号验证函数
	Validate.RegisterValidation("mobile", func(fl validator.FieldLevel) bool {
		matched, _ := regexp.MatchString(`^(1[1-9][0-9]\d{8})$`, fl.Field().String())
		return matched
	})

	// 自定义正整数验证函数
	Validate.RegisterValidation("positivenumber", func(fl validator.FieldLevel) bool {
		return fl.Field().Int() > 0
	})

	//return

	zh := zh.New()
	var found bool
	trans, found = ut.New(zh, zh).GetTranslator(zh.Locale())
	if !found {
		log.Fatal("translator not found")
	}
	zhTranslations.RegisterDefaultTranslations(Validate, trans)

	// 设置错误提示语
	Validate.RegisterTranslation("mobile", trans, func(ut ut.Translator) error {
		return ut.Add("mobile", "{0}格式不正确", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("mobile", fe.Field())
		return t
	})

	// 设置字段名名称
	Validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		tagName := field.Name
		label := field.Tag.Get("label")
		if label != "" {
			tagName = label
		}
		return tagName
	})
}

type ParamError struct {
	err string
}

func (e ParamError) Error() string {
	return e.err
}

func NewParamError(err string) ParamError {
	return ParamError{err: err}
}

type Validator interface {
	Validate() error
}

func (c *Context) Validate(data interface{}) error {

	if err := c.ShouldBind(data); err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			return NewParamError(err.Error())
		}
		validationErrorsTranslations := validationErrors.Translate(trans)
		errorText := ""
		for _, errText := range validationErrorsTranslations {
			errorText += errText + " "
		}
		return NewParamError(errorText)
	}

	if v, ok := data.(Validator); ok {
		if err := v.Validate(); err != nil {
			return NewParamError(err.Error())
		}
	}

	return nil
}
