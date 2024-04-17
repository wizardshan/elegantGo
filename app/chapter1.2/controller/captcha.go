package controller

import (
	"app/chapter1.2/controller/request"
	"app/chapter1.2/controller/response"
	"github.com/gin-gonic/gin"
)

type Captcha struct{}

func NewCaptcha() *Captcha {
	ctr := new(Captcha)
	return ctr
}

func (ctr *Captcha) Send(c *gin.Context) (response.Data, error) {

	request := new(request.CaptchaSend)
	if err := c.Validate(request); err != nil {
		return nil, err
	}

	return nil, nil
}
