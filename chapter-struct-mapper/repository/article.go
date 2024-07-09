package repository

import (
	"context"
	"elegantGo/chapter-struct-mapper/domain"
	"elegantGo/chapter-struct-mapper/repository/entity"
	"fmt"
	"time"
)

type Article struct {
}

func NewArticle() *Article {
	repo := new(Article)
	return repo
}

func (repo *Article) Fetch(ctx context.Context, id int) *domain.Article {
	entArticle := repo.entArticle(1)
	return entArticle.MapperCopier()
}

func (repo *Article) FetchMany(ctx context.Context) domain.Articles {

	entArticles := entity.Articles{
		repo.entArticle(1),
		repo.entArticle(2),
		repo.entArticle(3),
	}
	return entArticles.Mapper()
}

func (repo *Article) entArticle(i int) *entity.Article {
	entArticle := new(entity.Article)
	entArticle.ID = i
	entArticle.Title = fmt.Sprintf("title%d", i)
	entArticle.Content = fmt.Sprintf("content%d", i)
	entArticle.TimesOfRead = 100 * i
	entArticle.CreateTime = time.Now().AddDate(0, 0, i).UTC()
	return entArticle
}
