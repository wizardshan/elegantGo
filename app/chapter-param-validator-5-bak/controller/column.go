package controller

import (
	"app/chapter-param-validator-4/controller/request"
	"app/chapter-param-validator-4/controller/response"
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

func (ctr *Column) Detail(c *gin.Context) (response.Data, error) {
	request := new(request.ColumnDetail)
	if err := c.Validate(request); err != nil {
		return nil, err
	}

	return request, nil
}
