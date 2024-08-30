package controller

import (
	"elegantGo/chapter-param-complex-validator-2/controller/request"
	"elegantGo/chapter-param-complex-validator-2/controller/response"
	"fmt"
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

	fmt.Println(request.Amount.Values())
	fmt.Println(request.Level.Values())
	fmt.Println(request.Status.Values())
	fmt.Println(request.CreateTime.Values())
	fmt.Println(request.UpdateTime.Values())
	fmt.Println(request.StartTime.Values())

	return request, nil
}
