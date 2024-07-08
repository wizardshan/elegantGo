package main

import (
	"elegantGo/chapter-param-validator-xss/controller"
	"elegantGo/chapter-param-validator-xss/repository"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.New()
	engine.LoadHTMLGlob("chapter-param-validator-xss/controller/view/*.tmpl")
	repoArticle := repository.NewArticle()
	ctrArticle := controller.NewArticle(repoArticle)
	engine.GET("/article/search", ctrArticle.Search)

	engine.GET("/cookies", func(c *gin.Context) {
		fmt.Println(c.Request.URL.String())
	})

	engine.Run()
}
