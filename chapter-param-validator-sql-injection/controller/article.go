package controller

import (
	"elegantGo/chapter-param-validator-sql-injection/controller/response"
	"elegantGo/chapter-param-validator-sql-injection/repository"
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
	hashID := c.DefaultQuery("hashID", "")
	article := ctr.repo.Get(c.Request.Context(), hashID)

	return article, nil
}
