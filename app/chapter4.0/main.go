package main

import (
	"app/chapter4.0/controller"
	"app/chapter4.0/repository"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.New()

	handler := new(controller.Handler)

	repoArticle := repository.NewArticle()
	ctrArticle := controller.NewArticle(repoArticle)
	engine.GET("/article", handler.Wrapper(ctrArticle.Get))
	engine.GET("/articles", handler.Wrapper(ctrArticle.All))

	engine.Run()
}
