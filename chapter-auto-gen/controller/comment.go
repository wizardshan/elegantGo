package controller

import (
	"elegantGo/chapter-auto-gen/controller/request"
	"elegantGo/chapter-auto-gen/controller/response"
	"elegantGo/chapter-auto-gen/repository"
	"elegantGo/chapter-auto-gen/repository/ent"
	"github.com/gin-gonic/gin"
)

type Comment struct {
	repo *repository.Comment
}

func NewComment(repo *repository.Comment) *Comment {
	ctr := new(Comment)
	ctr.repo = repo
	return ctr
}

func (ctr *Comment) Many(c *gin.Context) (response.Data, error) {
	request := new(request.Comments)
	if err := c.ShouldBind(request); err != nil {
		return nil, err
	}

	domComments := ctr.repo.FetchMany(c.Request.Context(), func(builder *ent.CommentQuery) {
		builder.Offset(request.Offset.Value()).Limit(request.Limit.Value()).Order(request.Order.By(request.Sort.Value()))
	})

	var resp response.Comments
	return resp.Mapper(domComments), nil
}

func (ctr *Comment) One(c *gin.Context) (response.Data, error) {
	request := new(request.Comment)
	if err := c.ShouldBind(request); err != nil {
		return nil, err
	}

	domComment := ctr.repo.Fetch(c.Request.Context(), request.ID)

	var resp *response.Comment
	return resp.Mapper(domComment), nil
}
