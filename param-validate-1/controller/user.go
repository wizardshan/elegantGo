package controller

import (
	"elegantGo/param-validate-1/controller/response"
	"errors"
	"github.com/gin-gonic/gin"
	"regexp"
)

type User struct{}

func NewUser() *User {
	ctr := new(User)
	return ctr
}

func (ctr *User) Login(c *gin.Context) (response.Data, error) {
	mobile := c.DefaultQuery("mobile", "")
	captcha := c.DefaultQuery("captcha", "")

	if mobile == "" {
		return nil, errors.New("手机号不能为空")
	}

	if matched, _ := regexp.MatchString(`^[0-9]+$`, mobile); !matched {
		return nil, errors.New("手机号必须数字")
	}

	if matched, _ := regexp.MatchString(`^(1[3-9][0-9]\d{8})$`, mobile); !matched {
		return nil, errors.New("手机号格式不正确")
	}

	if captcha == "" {
		return nil, errors.New("验证码不能为空")
	}

	if matched, _ := regexp.MatchString(`^[0-9]+$`, captcha); !matched {
		return nil, errors.New("验证码必须数字")
	}

	if len(captcha) != 4 {
		return nil, errors.New("验证码必须4位")
	}

	return gin.H{"Mobile": mobile, "Captcha": captcha}, nil
}
