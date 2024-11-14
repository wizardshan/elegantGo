package main

import (
	"elegantGo/param-validate-xss/controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.New()
	engine.LoadHTMLGlob("param-validate-xss/controller/view/*.tmpl")
	ctrArticle := controller.NewArticle()
	engine.GET("/article/search", ctrArticle.Search)

	engine.GET("/cookies", func(c *gin.Context) {
		fmt.Println(c.Request.URL.String())
	})

	engine.Run()
}
