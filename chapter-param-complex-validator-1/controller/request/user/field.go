package user

type NicknameField struct {
	Nickname string `binding:"required,min=2,max=10"`
}

type PasswordField struct {
	Password string `binding:"required,min=6"`
}

type RePasswordField struct {
	RePassword string `binding:"eqfield=Password"`
}
