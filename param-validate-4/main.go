package main

import (
	"elegantGo/param-validate-4/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.New()
	ctrUser := controller.NewUser()
	engine.GET("/user/login", controller.Wrapper(ctrUser.Login))

	ctrSms := controller.NewSms()
	engine.GET("/sms/captcha", controller.Wrapper(ctrSms.Captcha))

	engine.Run()
}
