package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

type Captcha struct{}

func NewCaptcha() *Captcha {
	ctr := new(Captcha)
	return ctr
}

func (ctr *Captcha) Send(c *gin.Context) {
	mobile := c.DefaultQuery("mobile", "")

	if mobile == "" {
		c.JSON(http.StatusOK, gin.H{
			"error": "手机号不能为空",
		})
		return
	}

	matched, err := regexp.MatchString(`^(1[3-9][0-9]\d{8})$`, mobile)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		return
	}

	if !matched {
		c.JSON(http.StatusOK, gin.H{
			"error": "手机号格式不正确",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mobile": mobile,
	})
}
