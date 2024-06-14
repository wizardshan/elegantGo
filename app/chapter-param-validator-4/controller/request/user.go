package request

import (
	"app/chapter-param-validator-4/controller/request/user"
	"errors"
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

func (req *UserRegister) Validate() error {
	if req.Password != req.RePassword {
		return errors.New("两次密码不一致")
	}
	return nil
}

type UserMany struct {
	QueryField
	Filter struct {
		ID         *int
		Username   *string
		Gender     *NumbersBySeparatorField
		Status     *StringsBySeparatorField
		Level      *NumberRangeField
		CreateTime *TimeRangeField
	}
	Operator struct {
		ID         *EqualField
		Username   *LikeField
		Gender     *InField
		Status     *InField
		Level      *BetweenField
		CreateTime *BetweenField
	}
}

func (req *UserMany) Validate() error {
	return nil
}
