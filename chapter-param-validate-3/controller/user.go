package controller

import (
	"elegantGo/chapter-param-validate-3/controller/response"
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

	if empty(mobile) {
		return nil, ErrMobileEmpty
	}

	if !isNumber(mobile) {
		return nil, ErrMobileNotNumber
	}

	if !isMobile(mobile) {
		return nil, ErrMobileFormat
	}

	if empty(captcha) {
		return nil, ErrCaptchaEmpty
	}

	if !isNumber(captcha) {
		return nil, ErrCaptchaNotNumber
	}

	if !length(captcha, 4) {
		return nil, ErrCaptchaLength
	}

	return gin.H{"Mobile": mobile, "Captcha": captcha}, nil
}
