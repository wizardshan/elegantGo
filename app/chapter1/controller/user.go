package controller

import (
	"app/chapter1/controller/response"
	"github.com/gin-gonic/gin"
	"regexp"
)

type User struct{}

func NewUser() *User {
	ctr := new(User)
	return ctr
}

func (ctr *User) Login(c *gin.Context) (response.Data, error) {
	mobile := c.DefaultPostForm("mobile", "")
	captcha := c.DefaultPostForm("captcha", "")

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

	if captcha == "" {
		return nil, gin.NewParamError("验证码不能为空")
	}

	if len(captcha) != 4 {
		return nil, gin.NewParamError("验证码非法")
	}

	return nil, nil
}
