package request

type SmsCaptcha struct {
	Mobile string `binding:"required,number,mobile"`
}
