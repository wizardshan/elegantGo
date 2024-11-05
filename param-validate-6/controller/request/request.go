package request

type MobileField struct {
	Mobile string `binding:"number,mobile"`
}

type CaptchaField struct {
	Captcha string `binding:"number,len=4"`
}

type IDField struct {
	ID int `binding:"min=1"`
}
