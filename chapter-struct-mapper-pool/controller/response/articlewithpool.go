package response

import (
	"app/chapter-struct-mapper-pool/domain"
	"time"
)

type ArticlesWithPool []*ArticleWithPool

type ArticleWithPool struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	TimesOfRead int       `json:"timesOfRead"`
	CreateTime  time.Time `json:"createTime"`
}

func NewArticleWithPool() *ArticleWithPool {
	return poolArticle.Get()
}

func (respArticle *ArticleWithPool) Put() {
	poolArticle.Put(respArticle)
}

func (respArticle *ArticleWithPool) Mapper(domArticle *domain.Article) *ArticleWithPool {
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

func (respArticles ArticlesWithPool) Put() {
	for _, resp := range respArticles {
		resp.Put()
	}
}

func (respArticles ArticlesWithPool) Mapper(domArticles domain.Articles) ArticlesWithPool {

	size := len(domArticles)
	respArticles = make(ArticlesWithPool, size)
	for i := 0; i < size; i++ {
		respArticle := NewArticleWithPool()
		respArticles[i] = respArticle.Mapper(domArticles[i])
	}

	return respArticles
}
