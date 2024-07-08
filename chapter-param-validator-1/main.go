package main

import (
	"elegantGo/chapter-param-validator-1/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.New()
	ctrUser := controller.NewUser()
	engine.GET("/user/login", ctrUser.Login)

	ctrCaptcha := controller.NewCaptcha()
	engine.GET("/captcha/send", ctrCaptcha.Send)

	engine.Run()
}
