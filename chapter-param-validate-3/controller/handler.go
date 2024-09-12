package controller

import (
	"elegantGo/chapter-param-validate-3/controller/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerFunc func(c *gin.Context) (response.Data, error)

// 4xx client error you messed up 5xx server error I messed up
func Wrapper(handlerFunc HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatusJSON(http.StatusOK, response.Resp{Code: http.StatusInternalServerError, Message: fmt.Sprintf("%v", err)})
			}
		}()

		if data, err := handlerFunc(c); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, response.Resp{Code: http.StatusBadRequest, Message: err.Error()})
		} else {
			c.JSON(http.StatusOK, response.Resp{Code: http.StatusOK, Message: http.StatusText(http.StatusOK), Success: true, Data: data})
		}
	}
}
