package domain

import (
	"app/chapter5.0/pool"
)

var poolArticle = pool.New(func() *Article {
	return new(Article)
})
