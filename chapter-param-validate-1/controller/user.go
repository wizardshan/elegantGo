package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

type User struct{}

func NewUser() *User {
	ctr := new(User)
	return ctr
}

func (ctr *User) Login(c *gin.Context) {
	mobile := c.DefaultQuery("mobile", "")
	captcha := c.DefaultQuery("captcha", "")

	if mobile == "" {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": "手机号不能为空"})
		return
	}

	if matched, _ := regexp.MatchString(`^[0-9]+$`, mobile); !matched {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": "手机号必须数字"})
		return
	}

	if matched, _ := regexp.MatchString(`^(1[3-9][0-9]\d{8})$`, mobile); !matched {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": "手机号格式不正确"})
		return
	}

	if captcha == "" {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": "验证码不能为空"})
		return
	}

	if matched, _ := regexp.MatchString(`^[0-9]+$`, captcha); !matched {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": "验证码必须数字"})
		return
	}

	if len(captcha) != 4 {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": "验证码必须4位"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mobile":  mobile,
		"captcha": captcha,
	})
}
