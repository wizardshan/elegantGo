package user

type NicknameField struct {
	Nickname string `binding:"min=2,max=10"`
}

type PasswordField struct {
	Password string `binding:"min=6"`
}
