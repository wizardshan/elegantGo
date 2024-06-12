package request

type MobileField struct {
	Mobile string `form:"mobile" valid:"required~手机号不能为空,numeric~手机号码应该为数字型,stringlength(11|11)~手机号码长度错误,IsMobile~手机号码格式错误"`
}

type CaptchaField struct {
	Captcha string `form:"captcha" valid:"required~验证码不能为空,numeric~验证码应该为数字型,stringlength(4|4)~验证码长度错误"`
}

type IDField struct {
	ID int `form:"id" valid:"required~id不能为空,int~id应该为数字型"`
}
