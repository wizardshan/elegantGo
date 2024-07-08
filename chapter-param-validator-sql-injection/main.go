package main

import (
	"database/sql"
	"elegantGo/chapter-param-validator-sql-injection/controller"
	"elegantGo/chapter-param-validator-sql-injection/repository"
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
	engine.GET("/article", handler.Wrapper(ctrArticle.One))

	engine.Run()
}
