package request

import (
	"elegantGo/chapter-param-validator-4/controller/request/user"
)

type UserLogin struct {
	MobileField
	CaptchaField
}

type UserDelete struct {
	IDSField
}

//type UserRegister struct {
//	MobileField
//	CaptchaField
//	user.NicknameField
//	Password   string `binding:"required,min=6" label:"密码"`
//	RePassword string `binding:"eqfield=Password" label:"重复密码"`
//}

type UserRegister struct {
	MobileField
	CaptchaField
	user.NicknameField
	user.PasswordField
	user.RePasswordField
}
