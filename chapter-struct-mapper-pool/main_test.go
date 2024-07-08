package main

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

// go test -bench . -benchmem
func BenchmarkArticle(b *testing.B) {
	urlIndex := "/article"
	for n := 0; n < b.N; n++ {
		Get(urlIndex, engine)
	}
}

func BenchmarkArticleWithPool(b *testing.B) {
	urlIndex := "/articlePool"
	for n := 0; n < b.N; n++ {
		Get(urlIndex, engine)
	}
}

func BenchmarkArticles(b *testing.B) {
	urlIndex := "/articles"
	for n := 0; n < b.N; n++ {
		Get(urlIndex, engine)
	}
}

func BenchmarkArticlesWithPool(b *testing.B) {
	urlIndex := "/articlesPool"
	for n := 0; n < b.N; n++ {
		Get(urlIndex, engine)
	}
}

//func BenchmarkMapperWithCopier(b *testing.B) {
//	var entArticle Article
//	for n := 0; n < b.N; n++ {
//		entArticle.MapperWithCopier()
//	}
//}

func Get(uri string, router *gin.Engine) *httptest.ResponseRecorder {
	req := httptest.NewRequest("GET", uri, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
