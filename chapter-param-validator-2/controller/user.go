package controller

import (
	"elegantGo/chapter-param-validator-2/controller/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct{}

func NewUser() *User {
	ctr := new(User)
	return ctr
}

func (ctr *User) Login(c *gin.Context) {

	request := new(request.UserLogin)
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
