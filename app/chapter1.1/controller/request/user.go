package request

type UserLogin struct {
	Mobile string `form:"mobile" valid:"required~手机号不能为空,numeric~手机号码应该为数字型,IsMobile~手机号码格式错误"`
	Captcha string `form:"captcha" valid:"required~验证码不能为空,numeric~验证码应该为数字型"`
}