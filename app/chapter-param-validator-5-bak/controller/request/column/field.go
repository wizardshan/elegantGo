package column

type TitleField struct {
	Title string `form:"title" valid:"required~专栏标题不能为空,stringlength(4|50)~专栏标题4-50个字符"`
}
