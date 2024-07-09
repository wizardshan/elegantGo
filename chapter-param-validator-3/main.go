package main

import (
	"elegantGo/chapter-param-validator-3/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.New()
	handler := new(controller.Handler)

	ctrUser := controller.NewUser()
	engine.GET("/user/login", handler.Wrapper(ctrUser.Login))

	ctrCaptcha := controller.NewCaptcha()
	engine.GET("/captcha/send", handler.Wrapper(ctrCaptcha.Send))

	engine.Run()
}
