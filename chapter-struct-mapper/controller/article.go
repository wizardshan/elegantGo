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

	article := ctr.repo.Get(c.Request.Context(), request.ID)

	resp := new(response.Article)
	return resp.Mapper(article), nil
}

func (ctr *Article) Many(c *gin.Context) (response.Data, error) {

	articles := ctr.repo.Find(c.Request.Context())

	resp := response.Articles{}
	return resp.Mapper(articles), nil
}
