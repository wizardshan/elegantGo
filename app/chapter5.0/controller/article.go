package controller

import (
	"app/chapter5.0/controller/request"
	"app/chapter5.0/controller/response"
	"app/chapter5.0/repository"
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

	resp := new(response.Article)
	return resp.Mapper(article), nil
}

func (ctr *Article) GetWithPool(c *gin.Context) (response.Data, error) {

	request := request.NewArticleGetWithPool()
	if err := c.Validate(request); err != nil {
		return nil, err
	}
	defer request.Put()

	article := ctr.repo.GetWithPool(c.Request.Context(), request.ID)

	resp := response.NewArticleWithPool()
	return resp.Mapper(article), nil
}

func (ctr *Article) All(c *gin.Context) (response.Data, error) {

	request := new(request.ArticleAll)
	if err := c.Validate(request); err != nil {
		return nil, err
	}

	articles := ctr.repo.Find(c.Request.Context())

	resp := response.Articles{}
	return resp.Mapper(articles), nil
}

func (ctr *Article) AllWithPool(c *gin.Context) (response.Data, error) {

	request := request.NewArticleAllWithPool()
	if err := c.Validate(request); err != nil {
		return nil, err
	}
	defer request.Put()

	articles := ctr.repo.FindWithPool(c.Request.Context())

	resp := response.ArticlesWithPool{}
	return resp.Mapper(articles), nil
}
