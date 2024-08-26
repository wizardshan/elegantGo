package controller

import (
	"elegantGo/chapter-param-validate-4/controller/response"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type HandlerFunc func(c *gin.Context) (response.Data, error)

func Wrapper(handlerFunc HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {

		data, err := handlerFunc(c)
		if err == nil {
			c.JSON(http.StatusOK, response.Response{Code: response.OK, Message: response.Msg[response.OK], Data: data})
			return
		}

		code := response.BizError
		message := err.Error()
		var validErr validator.ValidationErrors
		var respErr response.Error
		if errors.As(err, &validErr) {
			code = response.ArgumentInvalid
		} else if errors.As(err, &respErr) {
			code = respErr.Code
		}

		c.AbortWithStatusJSON(http.StatusOK, response.Response{Code: code, Message: message, Data: data})
	}
}
