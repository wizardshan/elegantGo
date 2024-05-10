package response

type ArticleSearch struct {
	Keyword string `json:"keyword"`
}

func (resp *ArticleSearch) Mapper(keyword string) {
	resp.Keyword = keyword
}
