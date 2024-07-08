package request

type CaptchaSend struct {
	Mobile string `form:"mobile" binding:"required,number,mobile"`
}
