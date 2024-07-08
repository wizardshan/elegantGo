package response

import (
	"app/chapter-struct-mapper-pool/pool"
)

var poolArticle = pool.New(func() *ArticleWithPool {
	return new(ArticleWithPool)
})
