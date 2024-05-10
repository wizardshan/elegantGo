package main

import (
	"app/chapter3.0/controller"
	"app/chapter3.0/repository"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.New()
	engine.LoadHTMLGlob("controller/view/*.tmpl")
	repoArticle := repository.NewArticle()
	ctrArticle := controller.NewArticle(repoArticle)
	engine.GET("/article/search", ctrArticle.Search)

	engine.GET("/cookies", func(c *gin.Context) {
		fmt.Println(c.Request.URL.String())
	})

	engine.Run()
}
