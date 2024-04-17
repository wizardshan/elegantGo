package article

type TitleField struct {
	Title string `form:"title" valid:"required~文章标题不能为空,stringlength(4|50)~文章标题4-50个字符"`
}
