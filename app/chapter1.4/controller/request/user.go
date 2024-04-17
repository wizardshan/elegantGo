package request

import (
	"app/chapter1.4/controller/request/user"
	"errors"
)

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

func (req *UserRegister) Validate() error {
	if req.Password != req.RePassword {
		return errors.New("两次密码不一致")
	}

	return nil
}
