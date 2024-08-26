package main

import (
	"elegantGo/chapter-param-validate-1/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.New()
	ctrUser := controller.NewUser()
	engine.GET("/user/login", ctrUser.Login)

	ctrSms := controller.NewSms()
	engine.GET("/sms/captcha", ctrSms.Captcha)

	engine.Run()
}
