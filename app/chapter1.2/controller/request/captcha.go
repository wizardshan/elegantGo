package request

type CaptchaSend struct {
	Mobile string `form:"mobile" valid:"required~手机号不能为空,numeric~手机号码应该为数字型,IsMobile~手机号码格式错误"`
}
