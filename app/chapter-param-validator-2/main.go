package main

import (
	"app/chapter-param-validator-2/controller"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func main() {

	v := binding.Validator.Engine().(*validator.Validate)
	v.RegisterValidation("mobile", func(fl validator.FieldLevel) bool {
		matched, _ := regexp.MatchString(`^(1[1-9][0-9]\d{8})$`, fl.Field().String())
		return matched
	})

	engine := gin.New()

	ctrUser := controller.NewUser()
	engine.GET("/user/login", ctrUser.Login)

	ctrCaptcha := controller.NewCaptcha()
	engine.GET("/captcha/send", ctrCaptcha.Send)

	engine.Run()
}
