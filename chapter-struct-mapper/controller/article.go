package controller

import (
	"elegantGo/chapter-struct-mapper/controller/request"
	"elegantGo/chapter-struct-mapper/controller/response"
	"elegantGo/chapter-struct-mapper/repository"
	"github.com/gin-gonic/gin"
)

type Article struct {
	repo *repository.Article
}

func NewArticle(repo *repository.Article) *Article {
	ctr := new(Article)
	ctr.repo = repo
	return ctr
}

func (ctr *Article) One(c *gin.Context) (response.Data, error) {

	request := new(request.Article)
	if err := c.ShouldBind(request); err != nil {
		return nil, err
	}

	article := ctr.repo.Fetch(c.Request.Context(), request.ID)
	return response.MapperArticle(article), nil
}

func (ctr *Article) Many(c *gin.Context) (response.Data, error) {
	request := new(request.Articles)
	if err := c.ShouldBind(request); err != nil {
		return nil, err
	}

	articles := ctr.repo.FetchMany(c.Request.Context())
	return response.MapperArticles(articles), nil
}
