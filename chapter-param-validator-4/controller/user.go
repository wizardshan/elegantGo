package controller

import (
	"elegantGo/chapter-param-validator-4/controller/request"
	"elegantGo/chapter-param-validator-4/controller/response"
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

func (ctr *User) Delete(c *gin.Context) (response.Data, error) {
	request := new(request.UserDelete)
	if err := c.ShouldBind(request); err != nil {
		return nil, err
	}
	return request.GetIDS(), nil
}

func (ctr *User) Register(c *gin.Context) (response.Data, error) {
	request := new(request.UserRegister)
	if err := c.ShouldBind(request); err != nil {
		return nil, err
	}

	return request, nil
}
