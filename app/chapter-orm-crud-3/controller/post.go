package controller

import (
	"app/chapter-orm-crud-3/repository"
	"app/chapter-orm-crud-3/repository/ent"
	"app/chapter-orm-crud-3/repository/ent/comment"
	"app/chapter-orm-crud-3/repository/ent/post"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Post struct {
	repo        *repository.Post
	repoComment *repository.Comment
}

func NewPost(repo *repository.Post, repoComment *repository.Comment) *Post {
	ctr := new(Post)
	ctr.repo = repo
	ctr.repoComment = repoComment
	return ctr
}

func (ctr *Post) Many(c *gin.Context) {

	entPosts := ctr.repo.FetchMany(c.Request.Context(), func(builder *ent.PostQuery) {
		builder.WithUser()
	})
	c.JSON(http.StatusOK, entPosts)
}

func (ctr *Post) One(c *gin.Context) {
	id := 1
	posts := ctr.repo.FetchOne(c.Request.Context(), func(builder *ent.PostQuery) {
		builder.WithUser().WithComments(func(ops *ent.CommentQuery) {
			ops.WithUser()
		}).Where(post.ID(id))
	})
	c.JSON(http.StatusOK, posts)
}

func (ctr *Post) LatestComments(c *gin.Context) {
	comments := ctr.repoComment.FetchMany(c.Request.Context(), func(builder *ent.CommentQuery) {
		builder.WithUser().WithPost().Order(ent.Desc(comment.FieldCreateTime))
	})
	c.JSON(http.StatusOK, comments)
}
