package controller

import (
	"app/chapter-param-validator-5/controller/request"
	"app/chapter-param-validator-5/controller/response"
	"github.com/gin-gonic/gin"
)

type Column struct{}

func NewColumn() *Column {
	ctr := new(Column)
	return ctr
}

func (ctr *Column) Create(c *gin.Context) (response.Data, error) {
	request := new(request.ColumnCreate)
	if err := c.Validate(request); err != nil {
		return nil, err
	}

	return request, nil
}

func (ctr *Column) One(c *gin.Context) (response.Data, error) {
	request := new(request.ColumnOne)
	if err := c.Validate(request); err != nil {
		return nil, err
	}

	return request, nil
}
