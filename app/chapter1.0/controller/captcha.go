package controller

import (
	"app/chapter1.0/controller/response"
	"github.com/gin-gonic/gin"
	"regexp"
)

type Captcha struct{}

func NewCaptcha() *Captcha {
	ctr := new(Captcha)
	return ctr
}

func (ctr *Captcha) Send(c *gin.Context) (response.Data, error) {
	mobile := c.DefaultPostForm("mobile", "")

	if mobile == "" {
		return nil, gin.NewParamError("手机号不能为空")
	}

	matched, err := regexp.MatchString(`^(1[3-9][0-9]\d{8})$`, mobile)
	if err != nil {
		return nil, err
	}

	if !matched {
		return nil, gin.NewParamError("手机号格式不正确")
	}

	return nil, nil
}
