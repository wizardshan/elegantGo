package controller

import (
	"elegantGo/chapter-param-validator-2/controller/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Captcha struct{}

func NewCaptcha() *Captcha {
	ctr := new(Captcha)
	return ctr
}

func (ctr *Captcha) Send(c *gin.Context) {

	request := new(request.CaptchaSend)
	if err := c.ShouldBind(request); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"request": request,
	})
}
