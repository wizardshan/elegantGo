package controller

import (
	"elegantGo/chapter-param-validator-3/controller/request"
	"elegantGo/chapter-param-validator-3/controller/response"
	"github.com/gin-gonic/gin"
)

type Captcha struct{}

func NewCaptcha() *Captcha {
	ctr := new(Captcha)
	return ctr
}

func (ctr *Captcha) Send(c *gin.Context) (response.Data, error) {

	request := new(request.CaptchaSend)
	if err := c.ShouldBind(request); err != nil {
		return nil, err
	}

	return request, nil
}
