package controller

import (
	"elegantGo/chapter-orm-entgo/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Post struct {
	repo *repository.Post
}

func NewPost(repo *repository.Post) *Post {
	ctr := new(Post)
	ctr.repo = repo
	return ctr
}

func (ctr *Post) Many(c *gin.Context) {
	posts := ctr.repo.FetchMany(c.Request.Context())
	c.JSON(http.StatusOK, posts)
}

func (ctr *Post) One(c *gin.Context) {
	id := 1
	posts := ctr.repo.Fetch(c.Request.Context(), id)
	c.JSON(http.StatusOK, posts)
}

func (ctr *Post) Comments(c *gin.Context) {
	comments := ctr.repo.Comments(c.Request.Context())
	c.JSON(http.StatusOK, comments)
}
