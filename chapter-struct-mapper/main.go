package main

import (
	"elegantGo/chapter-struct-mapper/controller"
	"elegantGo/chapter-struct-mapper/repository"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.New()

	handler := new(controller.Handler)

	repoArticle := repository.NewArticle()
	ctrArticle := controller.NewArticle(repoArticle)
	engine.GET("/article", handler.Wrapper(ctrArticle.One))
	engine.GET("/articles", handler.Wrapper(ctrArticle.Many))

	engine.Run()
}
