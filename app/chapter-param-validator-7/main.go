package main

import (
	"app/chapter-param-validator-7/controller"
	"app/chapter-param-validator-7/repository"
	"app/chapter-param-validator-7/repository/ent"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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

	engine.Run()
}
