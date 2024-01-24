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

	r.GET("/user/exportV1", ctrUser.ExportV1)
	r.GET("/user/exportV2", ctrUser.ExportV2)

	r.Run()
}
