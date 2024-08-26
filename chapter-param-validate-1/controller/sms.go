package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

type Sms struct{}

func NewSms() *Sms {
	ctr := new(Sms)
	return ctr
}

func (ctr *Sms) Captcha(c *gin.Context) {
	mobile := c.DefaultQuery("mobile", "")

	if mobile == "" {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": "手机号不能为空"})
		return
	}

	if matched, _ := regexp.MatchString(`^(1[3-9][0-9]\d{8})$`, mobile); !matched {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": "手机号格式不正确"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mobile": mobile,
	})
}
