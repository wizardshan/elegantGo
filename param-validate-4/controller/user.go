package controller

import (
	"elegantGo/param-validate-4/controller/response"
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

	fields := Fields{
		{name: "手机号", funcNames: []string{NotEmpty, IsNumber, IsMobile}, value: mobile},
		{name: "验证码", funcNames: []string{NotEmpty, IsNumber, Length}, value: captcha, args: []any{4}},
	}

	if err := fields.validate(); err != nil {
		return nil, err
	}

	return gin.H{"Mobile": mobile, "Captcha": captcha}, nil
}
