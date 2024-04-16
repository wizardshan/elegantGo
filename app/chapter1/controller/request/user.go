package request

type UserEditProfile struct {
	Mobile   string `form:"mobile" binding:"required,min=11,max=11"`
	Nickname string `form:"nickname" binding:"required,min=2,max=10"`
	Bio      string `form:"bio"`
}
