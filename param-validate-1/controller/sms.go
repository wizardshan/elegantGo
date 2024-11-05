package controller

import (
	"elegantGo/param-validate-1/controller/response"
	"errors"
	"github.com/gin-gonic/gin"
	"regexp"
)

type Sms struct{}

func NewSms() *Sms {
	ctr := new(Sms)
	return ctr
}

func (ctr *Sms) Captcha(c *gin.Context) (response.Data, error) {
	mobile := c.DefaultQuery("mobile", "")

	if mobile == "" {
		return nil, errors.New("手机号不能为空")
	}

	if matched, _ := regexp.MatchString(`^[0-9]+$`, mobile); !matched {
		return nil, errors.New("手机号必须数字")
	}

	if matched, _ := regexp.MatchString(`^(1[3-9][0-9]\d{8})$`, mobile); !matched {
		return nil, errors.New("手机号格式不正确")
	}

	return gin.H{"Mobile": mobile}, nil
}
