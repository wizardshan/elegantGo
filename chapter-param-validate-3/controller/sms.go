package controller

import (
	"elegantGo/chapter-param-validate-3/controller/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Sms struct{}

func NewSms() *Sms {
	ctr := new(Sms)
	return ctr
}

func (ctr *Sms) Captcha(c *gin.Context) {
	request := new(request.SmsCaptcha)
	if err := c.ShouldBind(request); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"request": request,
	})
}
