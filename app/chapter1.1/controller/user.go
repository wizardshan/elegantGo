package controller

import (
	"app/chapter1.1/controller/request"
	"app/chapter1.1/controller/response"
	"github.com/asaskevich/govalidator"
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
		return nil, gin.NewParamError(err.Error())
	}

	if _, err := govalidator.ValidateStruct(request); err != nil {
		return nil, gin.NewParamError(err.Error())
	}

	return request, nil
}
