package request

import "elegantGo/chapter-param-validate-6/controller/request/user"

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
	// eqfield有效
	RePassword string `binding:"eqfield=Password"`
}
