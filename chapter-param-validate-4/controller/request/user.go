package request

import (
	"elegantGo/chapter-param-validate-4/controller/request/user"
)

type UserLogin struct {
	MobileField
	CaptchaField
}

type UserRegister struct {
	MobileField
	CaptchaField
	user.NicknameField
	user.PasswordField
	//user.RePasswordField
	RePassword string `binding:"eqfield=Password"`
}
