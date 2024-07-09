package controller

import (
	"elegantGo/chapter-struct-mapper/controller/response"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type HandlerFunc func(c *gin.Context) (response.Data, error)

type Handler struct {
}

func (handler *Handler) Wrapper(handlerFunc HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {

		data, err := handlerFunc(c)
		if err == nil {
			c.JSON(http.StatusOK, response.Response{Code: response.OK, Message: response.Msg[response.OK], Data: data, Success: true})
			return
		}

		code := response.BusinessError
		message := err.Error()
		var validErr validator.ValidationErrors
		if errors.As(err, &validErr) {
			code = response.InvalidArgument
		} else {
			var respErr response.Error
			if errors.As(err, &respErr) {
				code = respErr.Code
			}
		}

		c.JSON(http.StatusOK, response.Response{Code: code, Message: message, Data: data})
	}
}
