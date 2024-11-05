package user

type NicknameField struct {
	Nickname string `binding:"min=2,max=10"`
}

type PasswordField struct {
	Password string `binding:"min=6"`
}

// eqfield有效
type RePasswordField struct {
	//RePassword string `binding:"eqfield=Password"`
	//RePassword string `binding:"eqfield=PasswordField.Password"`
}
