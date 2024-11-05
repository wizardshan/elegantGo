package controller

import (
	"elegantGo/param-validate-4/controller/response"
	"github.com/gin-gonic/gin"
)

type Sms struct{}

func NewSms() *Sms {
	ctr := new(Sms)
	return ctr
}

func (ctr *Sms) Captcha(c *gin.Context) (response.Data, error) {
	mobile := c.DefaultQuery("mobile", "")

	fields := Fields{
		{name: "手机号", funcNames: []string{NotEmpty, IsNumber, IsMobile}, value: mobile},
	}

	if err := fields.validate(); err != nil {
		return nil, err
	}

	return gin.H{"Mobile": mobile}, nil
}
