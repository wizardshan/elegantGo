package main

import (
	"elegantGo/chapter-param-validator-4/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.New()
	handler := new(controller.Handler)

	ctrUser := controller.NewUser()
	engine.GET("/user/login", handler.Wrapper(ctrUser.Login))
	engine.GET("/user/delete", handler.Wrapper(ctrUser.Delete))
	engine.GET("/user/register", handler.Wrapper(ctrUser.Register))

	ctrCaptcha := controller.NewCaptcha()
	engine.GET("/captcha/send", handler.Wrapper(ctrCaptcha.Send))

	ctrColumn := controller.NewColumn()
	engine.GET("/column/create", handler.Wrapper(ctrColumn.Create))
	engine.GET("/column", handler.Wrapper(ctrColumn.One))

	engine.Run()
}
