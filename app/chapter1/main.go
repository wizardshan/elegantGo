package main

import (
	"app/chapter1/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()

	handler := new(controller.Handler)

	ctrUser := controller.NewUser()
	engine.POST("/user/login", handler.Wrapper(ctrUser.Login))

	ctrCaptcha := controller.NewCaptcha()
	engine.POST("/captcha/send", handler.Wrapper(ctrCaptcha.Send))

	engine.Run()
}
