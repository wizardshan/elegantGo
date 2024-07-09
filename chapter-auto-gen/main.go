package main

import (
	"elegantGo/chapter-auto-gen/controller"
	"elegantGo/chapter-auto-gen/repository"
	"elegantGo/chapter-auto-gen/repository/ent"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	host := "127.0.0.1:3306"
	name := "test"
	username := "root"
	password := ""

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true",
		username,
		password,
		host,
		name,
	)

	db, err := ent.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	engine := gin.New()

	handler := new(controller.Handler)

	repoUser := repository.NewUser(db)
	ctrUser := controller.NewUser(repoUser)
	engine.GET("/users", handler.Wrapper(ctrUser.Many))
	engine.GET("/user", handler.Wrapper(ctrUser.One))

	engine.Run()
}
