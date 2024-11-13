package controller

import (
	"elegantGo/param-validate-sql-injection/controller/request"
	"elegantGo/param-validate-sql-injection/controller/response"
	"elegantGo/param-validate-sql-injection/repository"
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

func (ctr *Article) Search(c *gin.Context) (response.Data, error) {
	request := new(request.ArticleSearch)
	if err := c.ShouldBind(request); err != nil {
		return nil, err
	}
	return request, nil
}

func (ctr *Article) One(c *gin.Context) (response.Data, error) {
	hashID := c.DefaultQuery("hashID", "")
	article := ctr.repo.Get(c.Request.Context(), hashID)

	return article, nil
}
