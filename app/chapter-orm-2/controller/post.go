package controller

import (
	"app/chapter-orm-2/repository"
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
	posts := ctr.repo.FetchByID(c.Request.Context(), id)
	posts.Comments = ctr.repo.Comments(c.Request.Context(), id)
	c.JSON(http.StatusOK, posts)
}
