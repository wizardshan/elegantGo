package controller

import (
	"elegantGo/param-validate-soc/controller/request"
	"elegantGo/param-validate-soc/controller/response"
	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() *User {
	ctr := new(User)
	return ctr
}

func (ctr *User) Many(c *gin.Context) (response.Data, error) {
	request := new(request.Users)
	if err := c.ShouldBind(request); err != nil {
		return nil, err
	}

	type V struct {
		Values       []int
		MustValues   []int
		ShouldValues []int
	}

	var err error
	v := new(V)
	v.Values = request.IDs.Values()
	v.MustValues, err = request.IDs.MustValues()
	v.ShouldValues = request.IDs.ShouldValues()

	return v, err
}
