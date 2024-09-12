package article

type TitleField struct {
	Title string `binding:"required,min=5,max=30"`
}
