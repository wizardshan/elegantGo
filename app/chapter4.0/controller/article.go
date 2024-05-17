package controller

import (
	"app/chapter4.0/controller/request"
	"app/chapter4.0/controller/response"
	"app/chapter4.0/repository"
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

func (ctr *Article) Get(c *gin.Context) (response.Data, error) {

	request := new(request.ArticleGet)
	if err := c.Validate(request); err != nil {
		return nil, err
	}

	article := ctr.repo.Get(c.Request.Context(), request.ID)

	resp := new(response.ArticleGet)
	return resp.Mapper(article), nil
}

func (ctr *Article) All(c *gin.Context) (response.Data, error) {

	articles := ctr.repo.Find(c.Request.Context())

	resp := response.ArticleAll{}
	return resp.Mapper(articles), nil
}
