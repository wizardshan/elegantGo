package domain

import (
	"elegantGo/chapter-struct-mapper-pool/pkg/pool"
	"time"
)

var poolArticle = pool.New(func() *Article {
	return new(Article)
})

type Articles []*Article

type Article struct {
	ID          int
	Title       string
	Content     string
	TimesOfRead int
	CreateTime  time.Time
}

func NewArticle() *Article {
	return poolArticle.Get()
}

func (domArticle *Article) Recycle() {
	poolArticle.Put(domArticle)
}

func (domArticles Articles) Recycle() {
	for _, domArticle := range domArticles {
		domArticle.Recycle()
	}
}
