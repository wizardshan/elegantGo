package controller

import (
	"elegantGo/chapter-param-validate-5/controller/request"
	"elegantGo/chapter-param-validate-5/controller/response"
	"github.com/gin-gonic/gin"
)

type Column struct{}

func NewColumn() *Column {
	ctr := new(Column)
	return ctr
}

func (ctr *Column) Create(c *gin.Context) (response.Data, error) {
	request := new(request.ColumnCreate)
	if err := c.ShouldBind(request); err != nil {
		return nil, err
	}

	return request, nil
}

func (ctr *Column) One(c *gin.Context) (response.Data, error) {
	request := new(request.ColumnOne)
	if err := c.ShouldBind(request); err != nil {
		return nil, err
	}

	return request, nil
}
