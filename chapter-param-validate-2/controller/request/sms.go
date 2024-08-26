package request

type SmsCaptcha struct {
	Mobile string `form:"mobile" binding:"required,number,mobile"`
}
