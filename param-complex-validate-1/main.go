package main

import (
	"elegantGo/param-complex-validate-1/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()

	ctrUser := controller.NewUser()
	engine.POST("/users", controller.Wrapper(ctrUser.Many))

	engine.Run()
}
