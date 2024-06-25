package controller

import (
	"app/chapter-param-complex-validator-1/controller/request"
	"app/chapter-param-complex-validator-1/controller/response"
	"github.com/gin-gonic/gin"
	"time"
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

	resp := struct {
		Sort   string
		IsDesc bool
		Offset int
		Limit  int

		Filter struct {
			ID         *int
			Nickname   *string
			CreateTime time.Time
		}
	}{
		Sort:   request.Sort.Value(),
		IsDesc: request.Order.IsDesc(),
		Offset: request.Offset.Value(),
		Limit:  request.Limit.Value(),
	}

	return resp, nil
}
