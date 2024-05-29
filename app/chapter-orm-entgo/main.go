package main

import (
	"app/chapter-orm-entgo/controller"
	"app/chapter-orm-entgo/repository"
	"app/chapter-orm-entgo/repository/ent"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	host := "127.0.0.1:3306"
	name := "test"
	username := "root"
	password := ""

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		name,
	)

	db, err := ent.Open("mysql", dsn)

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
