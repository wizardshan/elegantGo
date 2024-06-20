package controller

import (
	"app/chapter-param-validator-5/controller/request"
	"app/chapter-param-validator-5/controller/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct {
}

func NewUser() *User {
	ctr := new(User)
	return ctr
}

func (ctr *User) Many(c *gin.Context) (response.Data, error) {
	request := new(request.UserMany)
	if err := c.Validate(request); err != nil {
		return nil, err
	}

	fmt.Println("request.Page.Value=", request.Page.Value())

	return request, nil
}
