package response

import (
	"elegantGo/chapter-struct-mapper-pool/pool"
)

var poolArticle = pool.New(func() *ArticlePool {
	return new(ArticlePool)
})
