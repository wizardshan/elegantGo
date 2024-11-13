package main

import (
	"database/sql"
	"elegantGo/param-validate-sql-injection/controller"
	"elegantGo/param-validate-sql-injection/repository"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/test"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	engine := gin.New()

	repoArticle := repository.NewArticle(db)
	ctrArticle := controller.NewArticle(repoArticle)
	engine.GET("/article", controller.Wrapper(ctrArticle.One))
	engine.GET("/article/search", controller.Wrapper(ctrArticle.Search))

	engine.Run()
}
