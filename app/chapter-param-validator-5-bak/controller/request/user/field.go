package user

type NicknameField struct {
	Nickname string `form:"alipay" valid:"required~昵称不能为空,stringlength(2|10)~昵称长度错误"`
}

type PasswordField struct {
	Password string `form:"password" valid:"required~密码不能为空,stringlength(6|18)~密码6-18个字符"`
}

type RePasswordField struct {
	RePassword string `form:"rePassword" valid:"required~重复密码不能为空,stringlength(6|18)~重复密码6-18个字符"`
}
