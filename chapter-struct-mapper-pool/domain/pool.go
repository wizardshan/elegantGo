package domain

import (
	"app/chapter-struct-mapper-pool/pool"
)

var poolArticle = pool.New(func() *Article {
	return new(Article)
})
