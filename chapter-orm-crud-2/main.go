package main

import (
	"elegantGo/chapter-orm-crud-2/controller"
	"elegantGo/chapter-orm-crud-2/repository"
	"elegantGo/chapter-orm-crud-2/repository/ent"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// cd elegantGo/chapter-orm-crud-2/
// go generate ./repository/ent

func main() {

	host := "127.0.0.1:3306"
	name := "test"
	username := "root"
	password := ""

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true",
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
	repoUser := repository.NewUser(db)
	repoPost := repository.NewPost(db)
	repoComment := repository.NewComment(db)
	ctrPost := controller.NewPost(repoPost, repoComment)
	ctrUser := controller.NewUser(repoUser)
	engine.GET("/user", ctrUser.One)
	engine.GET("/users", ctrUser.Many)
	engine.GET("/user/register", ctrUser.Register)
	engine.GET("/user/upsert", ctrUser.Upsert)
	engine.GET("/user/rand", ctrUser.Rand)

	engine.GET("/posts", ctrPost.Many)
	engine.GET("/post", ctrPost.One)
	engine.GET("/post/comments", ctrPost.Comments)

	engine.Run()
}
