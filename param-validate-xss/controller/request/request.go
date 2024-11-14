package request

type MobileField struct {
	Mobile string `binding:"required,number,mobile"`
}

type CaptchaField struct {
	Captcha string `binding:"required,number,len=4"`
}

type IDField struct {
	ID int `binding:"required,min=1"`
}

type HashIDField struct {
	HashID string `binding:"required,alphanum"`
}

type KeywordField struct {
	//Keyword string `binding:"required"`
	Keyword string `binding:"required,xss"`
}
