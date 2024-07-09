package response

import (
	"elegantGo/chapter-struct-mapper-pool/domain"
	"elegantGo/chapter-struct-mapper-pool/pkg/pool"
)

var poolArticle = pool.New(func() *Article {
	return new(Article)
})

func NewArticle() *Article {
	return poolArticle.Get()
}

func (respArticle *Article) Recycle() {
	poolArticle.Put(respArticle)
}

func (respArticle *Article) MapperPool(domArticle *domain.Article) *Article {
	if domArticle == nil {
		return nil
	}
	defer domArticle.Recycle()

	respArticle.ID = domArticle.ID
	respArticle.Title = domArticle.Title
	respArticle.Content = domArticle.Content
	respArticle.TimesOfRead = domArticle.TimesOfRead
	respArticle.CreateTime = domArticle.CreateTime

	return respArticle
}

func (respArticles Articles) Recycle() {
	for _, resp := range respArticles {
		resp.Recycle()
	}
}

func (respArticles Articles) MapperPool(domArticles domain.Articles) Articles {

	defer respArticles.Recycle()

	size := len(domArticles)
	respArticles = make(Articles, size)
	for i := 0; i < size; i++ {
		respArticle := NewArticle()
		respArticles[i] = respArticle.Mapper(domArticles[i])
	}

	return respArticles
}
