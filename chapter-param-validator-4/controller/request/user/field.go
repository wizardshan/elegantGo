package user

type NicknameField struct {
	Nickname string `binding:"required,min=2,max=10" label:"昵称"`
}

type PasswordField struct {
	Password string `binding:"required,min=6" label:"密码"`
}
