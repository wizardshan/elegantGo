package controller

import (
	"app/chapter-param-validator-xss/controller/response"
	"errors"
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
			var paramError gin.Params
			if errors.As(err, &paramError) {
				code = response.ParamError
			}

			message = response.Msg[code] + "：" + err.Error()
		}
		c.JSON(http.StatusOK, response.Response{Code: code, Message: message, Data: data})
	}
}
