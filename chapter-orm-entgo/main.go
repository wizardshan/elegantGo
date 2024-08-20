package main

import (
	"elegantGo/chapter-orm-entgo/controller"
	"elegantGo/chapter-orm-entgo/repository"
	"elegantGo/chapter-orm-entgo/repository/ent"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/bbs?charset=utf8mb4&parseTime=true"
	db, err := ent.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	engine := gin.New()
	repoPost := repository.NewPost(db)
	ctrPost := controller.NewPost(repoPost)
	engine.GET("/posts", ctrPost.Many)
	engine.GET("/post", ctrPost.One)
	engine.GET("/post/comments", ctrPost.Comments)
	engine.Run()
}
