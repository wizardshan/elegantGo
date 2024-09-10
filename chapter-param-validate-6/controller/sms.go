package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Sms struct{}

func NewSms() *Sms {
	ctr := new(Sms)
	return ctr
}

func (ctr *Sms) Captcha(c *gin.Context) {
	mobile := c.DefaultQuery("mobile", "")

	fields := []*field{
		{name: "手机号", funcNames: []string{NotEmpty, IsMobile}, value: mobile},
	}

	if errs := validate(fields); errs != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": errs.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mobile": mobile,
	})
}
