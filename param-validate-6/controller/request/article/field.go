package article

type TitleField struct {
	Title string `binding:"min=5,max=30"`
}
