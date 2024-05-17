package repository

import (
	"app/chapter5.0/domain"
	"app/chapter5.0/repository/entity"
	"context"
	"time"
)

type Article struct {
}

func NewArticle() *Article {
	repo := new(Article)
	return repo
}

func (repo *Article) Get(ctx context.Context, id int) *domain.Article {

	entityArticle := new(entity.Article)
	entityArticle.ID = 1
	entityArticle.Title = "title1"
	entityArticle.Content = "content1"
	entityArticle.TimesOfRead = 100
	entityArticle.CreateTime = time.Now()
	return entityArticle.Mapper()
}

func (repo *Article) GetWithPool(ctx context.Context, id int) *domain.Article {

	entityArticle := new(entity.Article)
	entityArticle.ID = 1
	entityArticle.Title = "title1"
	entityArticle.Content = "content1"
	entityArticle.TimesOfRead = 100
	entityArticle.CreateTime = time.Now()
	return entityArticle.MapperWithPool()
}

func (repo *Article) Find(ctx context.Context) domain.Articles {

	entityArticle1 := new(entity.Article)
	entityArticle1.ID = 1
	entityArticle1.Title = "title1"
	entityArticle1.Content = "content1"
	entityArticle1.TimesOfRead = 100
	entityArticle1.CreateTime = time.Now()

	entityArticle2 := new(entity.Article)
	entityArticle2.ID = 2
	entityArticle2.Title = "title2"
	entityArticle2.Content = "content2"
	entityArticle2.TimesOfRead = 200
	entityArticle2.CreateTime = time.Now()

	var entityArticles entity.Articles
	entityArticles = append(entityArticles, entityArticle1, entityArticle2)
	//entityArticles = append(entityArticles, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2)

	return entityArticles.Mapper()
}

func (repo *Article) FindWithPool(ctx context.Context) domain.Articles {

	entityArticle1 := new(entity.Article)
	entityArticle1.ID = 1
	entityArticle1.Title = "title1"
	entityArticle1.Content = "content1"
	entityArticle1.TimesOfRead = 100
	entityArticle1.CreateTime = time.Now()

	entityArticle2 := new(entity.Article)
	entityArticle2.ID = 2
	entityArticle2.Title = "title2"
	entityArticle2.Content = "content2"
	entityArticle2.TimesOfRead = 200
	entityArticle2.CreateTime = time.Now()

	var entityArticles entity.Articles
	entityArticles = append(entityArticles, entityArticle1, entityArticle2)
	//entityArticles = append(entityArticles, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2, entityArticle1, entityArticle2)

	return entityArticles.MapperWithPool()
}
