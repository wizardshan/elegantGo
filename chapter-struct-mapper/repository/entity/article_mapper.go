package entity

import (
	"elegantGo/chapter-struct-mapper/domain"
	"github.com/elliotchance/pie/v2"
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

func (entArticles Articles) Mapper() domain.Articles {
	return pie.Map(entArticles, func(ent *Article) *domain.Article {
		return ent.Mapper()
	})
}

func (entArticle *Article) MapperCopier() *domain.Article {
	if entArticle == nil {
		return nil
	}

	domArticle := new(domain.Article)
	copier.Copy(domArticle, entArticle)
	return domArticle
}
