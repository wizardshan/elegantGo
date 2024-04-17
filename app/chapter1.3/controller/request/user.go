package request

import "app/chapter1.3/controller/request/user"

type UserLogin struct {
	MobileField
	CaptchaField
}

type UserRegister struct {
	MobileField
	CaptchaField
	user.PasswordField
	user.RePasswordField
	user.NicknameField
}
