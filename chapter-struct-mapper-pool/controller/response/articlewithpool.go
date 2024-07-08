package response

import (
	"elegantGo/chapter-struct-mapper-pool/domain"
	"time"
)

type ArticlesPool []*ArticlePool

type ArticlePool struct {
	ID          int
	Title       string
	Content     string
	TimesOfRead int
	CreateTime  time.Time
}

func NewArticlePool() *ArticlePool {
	return poolArticle.Get()
}

func (respArticle *ArticlePool) Put() {
	poolArticle.Put(respArticle)
}

func (respArticle *ArticlePool) Mapper(domArticle *domain.Article) *ArticlePool {
	if domArticle == nil {
		return nil
	}
	respArticle.ID = domArticle.ID
	respArticle.Title = domArticle.Title
	respArticle.Content = domArticle.Content
	respArticle.TimesOfRead = domArticle.TimesOfRead
	respArticle.CreateTime = domArticle.CreateTime
	domArticle.Put()

	return respArticle
}

func (respArticles ArticlesPool) Put() {
	for _, resp := range respArticles {
		resp.Put()
	}
}

func (respArticles ArticlesPool) Mapper(domArticles domain.Articles) ArticlesPool {

	size := len(domArticles)
	respArticles = make(ArticlesPool, size)
	for i := 0; i < size; i++ {
		respArticle := NewArticlePool()
		respArticles[i] = respArticle.Mapper(domArticles[i])
	}

	return respArticles
}
