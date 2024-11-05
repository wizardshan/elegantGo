package main

import (
	"elegantGo/param-validate-soc/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.New()
	ctrUser := controller.NewUser()
	engine.GET("/users", controller.Wrapper(ctrUser.Many))

	engine.Run()
}
