package request

type UserLogin struct {
	Mobile  string `binding:"required,number,mobile" label:"手机号码"`
	Captcha string `binding:"required,number,len=4" label:"手机验证码"`
}
