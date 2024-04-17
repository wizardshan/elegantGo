package controller

import (
	"app/chapter1.3/controller/request"
	"app/chapter1.3/controller/response"
	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() *User {
	ctr := new(User)
	return ctr
}

func (ctr *User) Login(c *gin.Context) (response.Data, error) {
	request := new(request.UserLogin)
	if err := c.Validate(request); err != nil {
		return nil, err
	}

	return request, nil
}

func (ctr *User) Register(c *gin.Context) (response.Data, error) {
	request := new(request.UserRegister)
	if err := c.Validate(request); err != nil {
		return nil, err
	}

	return request, nil
}
