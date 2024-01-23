package main

import (
	"elegantGo/lession01/controller"
	"elegantGo/lession01/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	repoUser := repository.NewUser()
	ctrUser := controller.NewUser(repoUser)

	r.GET("/user/export", ctrUser.Export)

	r.Run()
}
