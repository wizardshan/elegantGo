package controller

import (
	"elegantGo/chapter-param-validate-2/controller/response"
	"errors"
	"github.com/gin-gonic/gin"
)

type Sms struct{}

func NewSms() *Sms {
	ctr := new(Sms)
	return ctr
}

func (ctr *Sms) Captcha(c *gin.Context) (response.Data, error) {
	mobile := c.DefaultQuery("mobile", "")

	if !notEmpty(mobile) {
		return nil, errors.New("手机号不能为空")
	}

	if !isNumber(mobile) {
		return nil, errors.New("手机号必须数字")
	}

	if !isMobile(mobile) {
		return nil, errors.New("手机号格式不正确")
	}

	return gin.H{"Mobile": mobile}, nil
}
