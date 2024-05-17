package main

import (
	"app/chapter5.0/controller"
	"app/chapter5.0/repository"
	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func init() {
	engine = gin.New()
	handler := new(controller.Handler)

	repoArticle := repository.NewArticle()
	ctrArticle := controller.NewArticle(repoArticle)
	engine.GET("/article", handler.Wrapper(ctrArticle.Get))
	engine.GET("/articleWithPool", handler.Wrapper(ctrArticle.GetWithPool))
	engine.GET("/articles", handler.Wrapper(ctrArticle.All))
	engine.GET("/articlesWithPool", handler.Wrapper(ctrArticle.AllWithPool))
}
func main() {

	engine.Run()
}
