package request

type UserLogin struct {
	Mobile  string `form:"mobile" binding:"required,number,mobile"`
	Captcha string `form:"captcha" binding:"required,number,len=4"`
}
