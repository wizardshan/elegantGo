package main

import (
	"elegantGo/chapter-param-complex-validator-3/controller"
	"elegantGo/chapter-param-complex-validator-3/repository"
	"elegantGo/chapter-param-complex-validator-3/repository/ent"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	host := "127.0.0.1:3306"
	name := "test"
	username := "root"
	password := "123456"

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

	repoUser := repository.NewUser(db)
	ctrUser := controller.NewUser(repoUser)
	engine.GET("/users", controller.Wrapper(ctrUser.Many))

	engine.Run()
}
