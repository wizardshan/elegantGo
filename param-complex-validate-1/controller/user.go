package controller

import (
	"elegantGo/param-complex-validate-1/controller/request"
	"elegantGo/param-complex-validate-1/controller/response"
	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() *User {
	ctr := new(User)
	return ctr
}

func (ctr *User) Many(c *gin.Context) (response.Data, error) {
	req := new(request.UserMany)
	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}
	return req, nil
}
