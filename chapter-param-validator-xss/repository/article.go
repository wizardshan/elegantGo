package repository

type Article struct {
}

func NewArticle() *Article {
	repo := new(Article)
	return repo
}
