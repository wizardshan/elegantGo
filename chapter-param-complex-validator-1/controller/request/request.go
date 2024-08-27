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
