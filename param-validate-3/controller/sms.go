package controller

import (
	"elegantGo/param-validate-3/controller/response"
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
		return nil, ErrMobileNotEmpty
	}

	if !isNumber(mobile) {
		return nil, ErrMobileNotNumber
	}

	if !isMobile(mobile) {
		return nil, ErrMobileFormat
	}

	return gin.H{"Mobile": mobile}, nil
}
