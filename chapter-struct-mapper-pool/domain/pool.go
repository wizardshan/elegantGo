package domain

import (
	"elegantGo/chapter-struct-mapper-pool/pool"
)

var poolArticle = pool.New(func() *Article {
	return new(Article)
})
