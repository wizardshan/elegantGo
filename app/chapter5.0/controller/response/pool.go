package response

import (
	"app/chapter5.0/pool"
)

var poolArticle = pool.New(func() *ArticleWithPool {
	return new(ArticleWithPool)
})
