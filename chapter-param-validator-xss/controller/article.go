package controller

import (
	"app/chapter-param-validator-xss/controller/request"
	"app/chapter-param-validator-xss/repository"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

type Article struct {
	repo *repository.Article
}

func NewArticle(repo *repository.Article) *Article {
	ctr := new(Article)
	ctr.repo = repo
	return ctr
}

func (ctr *Article) Search(c *gin.Context) {

	request := new(request.ArticleSearch)
	if err := c.Validate(request); err != nil {
		panic(err)
	}

	// 种两个用于演示的cookie
	c.SetCookie("token", "token123456", 86400*30, "/", "127.0.0.1", false, false)
	c.SetCookie("userID", "1", 86400*30, "/", "127.0.0.1", false, false)

	c.HTML(http.StatusOK, "search.tmpl", gin.H{
		"keyword": template.HTML(request.Keyword), // 为了方便演示，template.HTML会显示原始字符串，默认会自动对特殊符号转义，
		//"keyword": request.Keyword,
	})
}
