package controller

import (
	"elegantGo/chapter-auto-gen/controller/request"
	"elegantGo/chapter-auto-gen/controller/response"
	"elegantGo/chapter-auto-gen/repository"
	"elegantGo/chapter-auto-gen/repository/ent"
	"elegantGo/chapter-auto-gen/repository/ent/post"
	"github.com/gin-gonic/gin"
)

type Post struct {
	repo *repository.Post
}

func NewPost(repo *repository.Post) *Post {
	ctr := new(Post)
	ctr.repo = repo
	return ctr
}

func (ctr *Post) Many(c *gin.Context) (response.Data, error) {
	request := new(request.Posts)
	if err := c.ShouldBind(request); err != nil {
		return nil, err
	}

	domPosts := ctr.repo.FetchMany(c.Request.Context(), func(builder *ent.PostQuery) {
		builder.WithUser().Offset(request.Offset.Value()).Limit(request.Limit.Value()).Order(request.Order.By(request.Sort.Value()))
	})

	var resp response.Posts
	return resp.Mapper(domPosts), nil
}

func (ctr *Post) One(c *gin.Context) (response.Data, error) {
	request := new(request.Post)
	if err := c.ShouldBind(request); err != nil {
		return nil, err
	}

	domPost := ctr.repo.FetchOne(c.Request.Context(), func(builder *ent.PostQuery) {
		builder.WithComments(func(query *ent.CommentQuery) {
			query.WithUser()
		}).WithUser().Where(post.ID(request.ID))
	})

	var resp *response.Post
	return resp.Mapper(domPost), nil
}
