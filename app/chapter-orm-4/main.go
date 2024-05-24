package main

import (
	"app/chapter-orm-4/controller"
	"app/chapter-orm-4/repository"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/test?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	engine := gin.New()
	repoPost := repository.NewPost(db)
	ctrPost := controller.NewPost(repoPost)
	engine.GET("/posts", ctrPost.Many)
	engine.GET("/post", ctrPost.One)
	engine.GET("/latestComments", ctrPost.LatestComments)
	engine.Run()
}
