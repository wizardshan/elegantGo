package controller

import (
	"elegantGo/param-validate-5/controller/request"
	"elegantGo/param-validate-5/controller/response"
	"github.com/gin-gonic/gin"
)

type Sms struct{}

func NewSms() *Sms {
	ctr := new(Sms)
	return ctr
}

func (ctr *Sms) Captcha(c *gin.Context) (response.Data, error) {
	request := new(request.SmsCaptcha)
	if err := c.ShouldBind(request); err != nil {
		return nil, err
	}

	return request, nil
}
