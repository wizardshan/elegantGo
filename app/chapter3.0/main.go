package main

import (
	"app/chapter3.0/controller"
	"app/chapter3.0/repository"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/test"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	//handler := new(controller.Handler)

	engine := gin.New()
	engine.LoadHTMLGlob("controller/view/*.tmpl")
	repoArticle := repository.NewArticle(db)
	ctrArticle := controller.NewArticle(repoArticle)
	engine.GET("/article/search", ctrArticle.Search)

	engine.GET("/cookies", func(c *gin.Context) {
		fmt.Println(c.Request.URL.String())
	})

	engine.Run()
}
