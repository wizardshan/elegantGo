package main

import (
	"app/chapter-param-validator-1/controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	ss := []int{1, 10}
	err := gin.Validate.Var(ss, "")
	fmt.Println(err)
	return

	engine := gin.New()
	ctrUser := controller.NewUser()
	engine.GET("/user/login", ctrUser.Login)

	ctrCaptcha := controller.NewCaptcha()
	engine.GET("/captcha/send", ctrCaptcha.Send)

	engine.Run()
}
