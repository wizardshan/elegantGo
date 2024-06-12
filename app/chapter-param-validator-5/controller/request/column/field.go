package column

type TitleField struct {
	Title string `binding:"required,min=4,max=50" label:"标题"`
}
