package request

type CaptchaSend struct {
	Mobile string `binding:"required,number,mobile"`
}
