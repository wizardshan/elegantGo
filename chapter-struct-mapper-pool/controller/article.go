package controller

import (
	"elegantGo/chapter-struct-mapper-pool/controller/request"
	"elegantGo/chapter-struct-mapper-pool/controller/response"
	"elegantGo/chapter-struct-mapper-pool/repository"
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

	resp := new(response.Article)
	return resp.Mapper(article), nil
}

func (ctr *Article) OnePool(c *gin.Context) (response.Data, error) {

	request := request.NewArticle()
	defer request.Recycle()
	if err := c.ShouldBind(request); err != nil {
		return nil, err
	}

	article := ctr.repo.FetchPool(c.Request.Context(), request.ID)

	resp := response.NewArticle()
	return resp.Mapper(article), nil
}

func (ctr *Article) Many(c *gin.Context) (response.Data, error) {

	request := new(request.Articles)
	if err := c.ShouldBind(request); err != nil {
		return nil, err
	}

	articles := ctr.repo.FetchMany(c.Request.Context())

	resp := response.Articles{}
	return resp.Mapper(articles), nil
}

func (ctr *Article) ManyPool(c *gin.Context) (response.Data, error) {

	request := request.NewArticles()
	defer request.Recycle()
	if err := c.ShouldBind(request); err != nil {
		return nil, err
	}

	articles := ctr.repo.FetchManyPool(c.Request.Context())

	resp := response.Articles{}
	return resp.MapperPool(articles), nil
}
