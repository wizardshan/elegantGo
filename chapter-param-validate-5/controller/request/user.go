package request

type UserLogin struct {
	Mobile  string `binding:"required,number,mobile"`
	Captcha string `binding:"required,number,len=4"`
}
