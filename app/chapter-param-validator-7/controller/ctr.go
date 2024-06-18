package controller

import (
	"app/chapter-param-validator-7/controller/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerFunc func(c *gin.Context) (response.Data, error)

type Handler struct {
}

func (handler *Handler) Wrapper(handlerFunc HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {

		data, err := handlerFunc(c)
		code := response.OK
		message := response.Msg[response.OK]

		if err != nil {
			code = response.InternalError
			if _, ok := err.(gin.ParamError); ok {
				code = response.ParamError
			} else if responseError, ok := err.(response.Error); ok {
				code = responseError.Code
			}

			message = err.Error()
			if msg, ok := response.Msg[code]; ok {
				message = msg + "：" + message
			}
		}
		c.JSON(http.StatusOK, response.Response{Code: code, Message: message, Data: data})
	}
}
