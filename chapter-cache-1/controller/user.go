package controller

import (
	"elegantGo/chapter-cache-1/controller/request"
	"elegantGo/chapter-cache-1/controller/response"
	"elegantGo/chapter-cache-1/repository"
	"elegantGo/chapter-cache-1/repository/ent"
	"github.com/gin-gonic/gin"
)

type User struct {
	repo *repository.User
}

func NewUser(repo *repository.User) *User {
	ctr := new(User)
	ctr.repo = repo
	return ctr
}

func (ctr *User) Many(c *gin.Context) (response.Data, error) {
	request := new(request.Users)
	if err := c.ShouldBind(request); err != nil {
		return nil, err
	}

	domUsers := ctr.repo.FetchMany(c.Request.Context(), func(builder *ent.UserQuery) {
		builder.Offset(request.Offset.Value()).Limit(request.Limit.Value()).Order(request.Order.By(request.Sort.Value()))
	})

	var resp response.Users
	return resp.Mapper(domUsers), nil
}

func (ctr *User) One(c *gin.Context) (response.Data, error) {
	request := new(request.User)
	if err := c.ShouldBind(request); err != nil {
		return nil, err
	}

	domUser := ctr.repo.Fetch(c.Request.Context(), request.ID)

	var resp *response.User
	return resp.Mapper(domUser), nil
}
