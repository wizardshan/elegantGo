package main

import (
	"app/chapter-param-complex-validator-2/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.New()
	handler := new(controller.Handler)

	ctrUser := controller.NewUser()
	engine.GET("/users", handler.Wrapper(ctrUser.Many))

	engine.Run()
}
