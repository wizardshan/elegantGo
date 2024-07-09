package controller

import (
	"elegantGo/chapter-orm-crud-2/repository"
	"elegantGo/chapter-orm-crud-2/repository/ent"
	"elegantGo/chapter-orm-crud-2/repository/ent/comment"
	"elegantGo/chapter-orm-crud-2/repository/ent/post"
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

func (ctr *Post) Comments(c *gin.Context) {
	comments := ctr.repoComment.FetchMany(c.Request.Context(), func(builder *ent.CommentQuery) {
		builder.WithUser().WithPost().Order(ent.Desc(comment.FieldCreateTime))
	})
	c.JSON(http.StatusOK, comments)
}
