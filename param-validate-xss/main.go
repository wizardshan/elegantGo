package main

import (
	"elegantGo/param-validate-xss/controller"
	"elegantGo/param-validate-xss/repository"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.New()
	engine.LoadHTMLGlob("param-validate-xss/controller/view/*.tmpl")
	repoArticle := repository.NewArticle()
	ctrArticle := controller.NewArticle(repoArticle)
	engine.GET("/article/search", ctrArticle.Search)

	engine.GET("/cookies", func(c *gin.Context) {
		fmt.Println(c.Request.URL.String())
	})

	engine.Run()
}
