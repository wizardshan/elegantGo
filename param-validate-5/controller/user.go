package controller

import (
	"elegantGo/param-validate-5/controller/request"
	"elegantGo/param-validate-5/controller/response"
	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() *User {
	ctr := new(User)
	return ctr
}

func (ctr *User) Login(c *gin.Context) (response.Data, error) {
	request := new(request.UserLogin)
	if err := c.ShouldBind(request); err != nil {
		return nil, err
	}

	return request, nil
}
