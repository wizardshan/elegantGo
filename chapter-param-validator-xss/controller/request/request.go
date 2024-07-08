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

type HashIDField struct {
	HashID int `form:"hashID" valid:"required~hashID不能为空,alphanum~hashID必须为字母和数字,stringlength(8|15)~hashID长度错误"`
}

type KeywordField struct {
	Keyword string `form:"keyword" valid:"required~搜索关键词不能为空"`
	//Keyword string `form:"keyword" valid:"required~搜索关键词不能为空,CheckXSS~非法字符"`
}
