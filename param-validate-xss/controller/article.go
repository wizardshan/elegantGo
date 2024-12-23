package controller

import (
	"elegantGo/param-validate-xss/controller/request"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

type Article struct {
}

func NewArticle() *Article {
	ctr := new(Article)
	return ctr
}

func (ctr *Article) Search(c *gin.Context) {

	request := new(request.ArticleSearch)
	if err := c.ShouldBind(request); err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
		return
	}

	// 种两个用于演示的cookie
	c.SetCookie("token", "token123456", 86400*30, "/", "127.0.0.1", false, false)
	c.SetCookie("userID", "1", 86400*30, "/", "127.0.0.1", false, false)

	c.HTML(http.StatusOK, "search.tmpl", gin.H{
		"keyword": template.HTML(request.Keyword), // 为了方便演示，template.HTML会显示原始字符串
		//"keyword": request.Keyword, // 默认会自动对特殊符号转义
	})
}
