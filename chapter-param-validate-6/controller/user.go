package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct{}

func NewUser() *User {
	ctr := new(User)
	return ctr
}

func (ctr *User) Login(c *gin.Context) {
	mobile := c.DefaultQuery("mobile", "")
	captcha := c.DefaultQuery("captcha", "")

	fields := []*field{
		{name: "手机号", funcNames: []string{NotEmpty, IsMobile}, value: mobile},
		{name: "验证码", funcNames: []string{NotEmpty, Length, IsNumber}, value: captcha, args: []any{4}},
	}

	if errs := validate(fields); errs != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": errs.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mobile":  mobile,
		"captcha": captcha,
	})
}
