package controller

import (
	"elegantGo/chapter-param-validate-2/controller/response"
	"errors"
	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() *User {
	ctr := new(User)
	return ctr
}

func (ctr *User) Login(c *gin.Context) (response.Data, error) {
	mobile := c.DefaultQuery("mobile", "")
	captcha := c.DefaultQuery("captcha", "")

	if !notEmpty(mobile) {
		return nil, errors.New("手机号不能为空")
	}

	if !isNumber(mobile) {
		return nil, errors.New("手机号必须数字")
	}

	if !isMobile(mobile) {
		return nil, errors.New("手机号格式不正确")
	}

	if !notEmpty(captcha) {
		return nil, errors.New("验证码不能为空")
	}

	if !isNumber(captcha) {
		return nil, errors.New("验证码必须数字")
	}

	if !length(captcha, 4) {
		return nil, errors.New("验证码必须4位")
	}

	return gin.H{"Mobile": mobile, "Captcha": captcha}, nil
}
