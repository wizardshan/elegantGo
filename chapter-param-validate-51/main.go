package main

import (
	"elegantGo/chapter-param-validate-5/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.New()

	ctrUser := controller.NewUser()
	engine.GET("/user/login", controller.Wrapper(ctrUser.Login))
	engine.GET("/user/register", controller.Wrapper(ctrUser.Register))

	ctrCaptcha := controller.NewCaptcha()
	engine.GET("/captcha/send", controller.Wrapper(ctrCaptcha.Send))

	ctrColumn := controller.NewColumn()
	engine.GET("/column/create", controller.Wrapper(ctrColumn.Create))
	engine.GET("/column", controller.Wrapper(ctrColumn.One))

	engine.Run()
}
