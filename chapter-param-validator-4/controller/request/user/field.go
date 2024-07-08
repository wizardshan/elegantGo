package user

type NicknameField struct {
	Nickname string `binding:"required,min=2,max=10" label:"昵称"`
}

type PasswordField struct {
	Password string `binding:"required,min=6" label:"密码"`
}

type RePasswordField struct {
	RePassword string `binding:"required" label:"重复密码"`
	//RePassword string `binding:"eqfield=Password" label:"重复密码"`
}
