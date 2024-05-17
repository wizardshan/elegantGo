package response

import (
	"app/chapter4.0/domain"
	"time"
)

type Articles []*Article

type Article struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	TimesOfRead int       `json:"timesOfRead"`
	CreateTime  time.Time `json:"createTime"`
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
	TimesOfRead Omit `json:"timesOfRead,omitempty"`
	//TimesOfRead int `json:"timesOfRead,omitempty"`
}

func (resp *ArticleOmit) Mapper(domainArticle *domain.Article) *ArticleOmit {
	if domainArticle == nil {
		return nil
	}

	resp.Article.Mapper(domainArticle)

	return resp
}
