package main

import (
	"elegantGo/chapter-param-complex-validator-1/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.New()
	handler := new(controller.Handler)

	ctrUser := controller.NewUser()
	engine.GET("/users", handler.Wrapper(ctrUser.Many))

	engine.Run()
}
