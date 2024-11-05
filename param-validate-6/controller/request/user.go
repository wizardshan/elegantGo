package request

import "elegantGo/param-validate-6/controller/request/user"

type UserLogin struct {
	MobileField
	CaptchaField
}

type UserRegister struct {
	MobileField
	CaptchaField
	user.NicknameField
	user.PasswordField
	RePassword string `binding:"eqfield=Password"`
}

type UserMany struct {
	MobileField        `binding:"omitempty"`
	user.NicknameField `binding:"omitempty"`
}
