package entity

import (
	"elegantGo/chapter-struct-mapper-pool/domain"
	"github.com/jinzhu/copier"
)

func (entArticle *Article) Mapper() *domain.Article {
	if entArticle == nil {
		return nil
	}

	domArticle := new(domain.Article)
	domArticle.ID = entArticle.ID
	domArticle.Title = entArticle.Title
	domArticle.Content = entArticle.Content
	domArticle.TimesOfRead = entArticle.TimesOfRead
	domArticle.CreateTime = entArticle.CreateTime

	return domArticle
}

func (entArticle *Article) MapperPool() *domain.Article {
	if entArticle == nil {
		return nil
	}

	domArticle := domain.NewArticle()
	domArticle.ID = entArticle.ID
	domArticle.Title = entArticle.Title
	domArticle.Content = entArticle.Content
	domArticle.TimesOfRead = entArticle.TimesOfRead
	domArticle.CreateTime = entArticle.CreateTime

	return domArticle
}

func (entArticle *Article) MapperCopier() *domain.Article {
	if entArticle == nil {
		return nil
	}
	domArticle := new(domain.Article)
	copier.Copy(entArticle, domArticle)

	return domArticle
}

func (entArticles Articles) Mapper() domain.Articles {

	size := len(entArticles)
	domArticles := make(domain.Articles, size)

	for i := 0; i < size; i++ {
		domArticles[i] = entArticles[i].Mapper()
	}

	return domArticles
}

func (entArticles Articles) MapperPool() domain.Articles {

	size := len(entArticles)
	domArticles := make(domain.Articles, size)

	for i := 0; i < size; i++ {
		domArticles[i] = entArticles[i].MapperPool()
	}

	return domArticles
}
