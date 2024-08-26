package controller

import (
	"elegantGo/chapter-param-validate-3/controller/request"
	"elegantGo/chapter-param-validate-3/controller/response"
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
