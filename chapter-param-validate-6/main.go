package main

import (
	"elegantGo/chapter-param-validate-6/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.New()
	ctrUser := controller.NewUser()
	engine.GET("/user/login", controller.Wrapper(ctrUser.Login))
	engine.GET("/user/register", controller.Wrapper(ctrUser.Register))

	ctrSms := controller.NewSms()
	engine.GET("/sms/captcha", controller.Wrapper(ctrSms.Captcha))

	engine.Run()
}
