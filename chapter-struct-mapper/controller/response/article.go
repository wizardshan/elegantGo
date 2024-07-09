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

type ArticleOmit struct {
	Article
	TimesOfRead Omit `json:"omitempty"`
	//TimesOfRead int `json:"omitempty"`
}

func MapperArticle(domArticle *domain.Article) *Article {
	if domArticle == nil {
		return nil
	}
	respArticle := new(Article)
	respArticle.ID = domArticle.ID
	respArticle.Title = domArticle.Title
	respArticle.Content = domArticle.Content
	respArticle.TimesOfRead = domArticle.TimesOfRead
	respArticle.CreateTime = domArticle.CreateTime
	return respArticle
}

func MapperArticles(domainArticles domain.Articles) Articles {
	return mapper(domainArticles, MapperArticle)
}
