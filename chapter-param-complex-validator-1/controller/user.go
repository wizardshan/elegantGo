package controller

import (
	"elegantGo/chapter-param-complex-validator-1/controller/request"
	"elegantGo/chapter-param-complex-validator-1/controller/response"
	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() *User {
	ctr := new(User)
	return ctr
}

func (ctr *User) Many(c *gin.Context) (response.Data, error) {
	request := new(request.UserMany)
	if err := c.ShouldBind(request); err != nil {
		return nil, err
	}

	return request.IDs.Values(), nil
}
