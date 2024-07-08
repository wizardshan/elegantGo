package main

import (
	"elegantGo/chapter-struct-mapper-pool/controller"
	"elegantGo/chapter-struct-mapper-pool/repository"
	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func init() {
	engine = gin.New()
	handler := new(controller.Handler)

	repoArticle := repository.NewArticle()
	ctrArticle := controller.NewArticle(repoArticle)
	engine.GET("/article", handler.Wrapper(ctrArticle.One))
	engine.GET("/articlePool", handler.Wrapper(ctrArticle.OnePool))
	engine.GET("/articles", handler.Wrapper(ctrArticle.Many))
	engine.GET("/articlesPool", handler.Wrapper(ctrArticle.ManyPool))
}

func main() {
	engine.Run()
}
