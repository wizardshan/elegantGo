package main

import (
	"elegantGo/chapter-orm-crud-1/controller"
	"elegantGo/chapter-orm-crud-1/repository"
	"elegantGo/chapter-orm-crud-1/repository/ent"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// cd chapter-orm-crud-1/
// go generate ./repository/ent

func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/bbs?charset=utf8mb4&parseTime=true"
	db, err := ent.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	engine := gin.New()
	repoUser := repository.NewUser(db)
	repoPost := repository.NewPost(db)
	repoComment := repository.NewComment(db)
	ctrPost := controller.NewPost(repoPost, repoComment)
	ctrUser := controller.NewUser(repoUser)
	engine.GET("/user", ctrUser.One)
	engine.GET("/users", ctrUser.Many)
	engine.GET("/user/register", ctrUser.Register)

	engine.GET("/posts", ctrPost.Many)
	engine.GET("/post", ctrPost.One)
	engine.GET("/post/comments", ctrPost.Comments)
	engine.Run()
}
