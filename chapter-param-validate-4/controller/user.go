package controller

import (
	"elegantGo/chapter-param-validate-4/controller/response"
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

	fields := []*Field{
		{name: "手机号", funcs: []string{NotEmpty, IsNumber, IsMobile}, value: mobile},
		{name: "验证码", funcs: []string{NotEmpty, IsNumber, Length}, value: captcha, args: []any{4}},
	}

	if err := validate(fields); err != nil {
		return nil, err
	}

	return gin.H{"Mobile": mobile, "Captcha": captcha}, nil
}
