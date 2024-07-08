package response

import (
	"elegantGo/chapter-struct-mapper/domain"
	"time"
)

type Articles []*Article

type Article struct {
	ID          int
	Title       string
	Content     string
	TimesOfRead int
	CreateTime  time.Time
}

func (respArticle *Article) Mapper(domArticle *domain.Article) *Article {
	if domArticle == nil {
		return nil
	}
	respArticle.ID = domArticle.ID
	respArticle.Title = domArticle.Title
	respArticle.Content = domArticle.Content
	respArticle.TimesOfRead = domArticle.TimesOfRead
	respArticle.CreateTime = domArticle.CreateTime

	return respArticle
}

func (respArticles Articles) Mapper(domArticles domain.Articles) Articles {

	size := len(domArticles)
	respArticles = make(Articles, size)
	for i := 0; i < size; i++ {
		var respArticle Article
		respArticles[i] = respArticle.Mapper(domArticles[i])
	}

	return respArticles
}

type ArticleGet struct {
	Article
}

type ArticleAll struct {
	Articles
}

type ArticleOmit struct {
	Article
	TimesOfRead Omit `json:"omitempty"`
	//TimesOfRead int `json:"omitempty"`
}

func (resp *ArticleOmit) Mapper(domainArticle *domain.Article) *ArticleOmit {
	if domainArticle == nil {
		return nil
	}

	resp.Article.Mapper(domainArticle)

	return resp
}
