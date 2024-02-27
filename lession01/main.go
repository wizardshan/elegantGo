package main

import (
	"elegantGo/lession01/controller"
	"elegantGo/lession01/repository"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	repoUser := repository.NewUser()
	ctrUserV1 := controller.NewUserV1(repoUser)
	ctrUserV2 := controller.NewUserV2(repoUser)
	ctrUserV3 := controller.NewUserV3(repoUser)

	r.GET("/userV1/export", ctrUserV1.Export)
	r.GET("/userV2/export", ctrUserV2.Export)
	r.GET("/userV3/export", ctrUserV3.Export)

	r.Run()
}
