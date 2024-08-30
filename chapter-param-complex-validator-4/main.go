package main

import (
	"elegantGo/chapter-param-complex-validator-2/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.New()

	ctrUser := controller.NewUser()
	engine.GET("/users", controller.Wrapper(ctrUser.Many))

	engine.Run()
}
