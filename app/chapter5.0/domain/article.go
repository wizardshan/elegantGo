package domain

import (
	"time"
)

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

func (domArticle *Article) Put() {
	poolArticle.Put(domArticle)
}
