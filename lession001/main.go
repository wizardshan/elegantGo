package main

import (
	"elegantGo/lession02/controller"
	"elegantGo/lession02/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	repoUser := repository.NewUser()
	ctrUser := controller.NewUser(repoUser)

	r.GET("/user/export", ctrUser.Export)

	r.Run()
}
