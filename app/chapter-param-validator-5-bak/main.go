package main

import (
	"app/chapter-param-validator-4/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.New()

	handler := new(controller.Handler)

	ctrUser := controller.NewUser()
	engine.POST("/user/login", handler.Wrapper(ctrUser.Login))
	engine.POST("/user/register", handler.Wrapper(ctrUser.Register))

	ctrCaptcha := controller.NewCaptcha()
	engine.POST("/captcha/send", handler.Wrapper(ctrCaptcha.Send))

	ctrColumn := controller.NewColumn()
	engine.POST("/column/create", handler.Wrapper(ctrColumn.Create))
	engine.GET("/column", handler.Wrapper(ctrColumn.Detail))

	engine.Run()
}
