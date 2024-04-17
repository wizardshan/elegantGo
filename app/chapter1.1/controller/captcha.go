package controller

import (
	"app/chapter1.1/controller/request"
	"app/chapter1.1/controller/response"
	"github.com/asaskevich/govalidator"
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
		return nil, gin.NewParamError(err.Error())
	}

	if _, err := govalidator.ValidateStruct(request); err != nil {
		return nil, gin.NewParamError(err.Error())
	}

	return nil, nil
}
