package main

import (
	"app/chapter-param-validator-sql-injection/controller"
	"app/chapter-param-validator-sql-injection/repository"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/test"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	engine := gin.New()

	handler := new(controller.Handler)

	repoArticle := repository.NewArticle(db)
	ctrArticle := controller.NewArticle(repoArticle)
	engine.GET("/article", handler.Wrapper(ctrArticle.Detail))

	engine.Run()
}
