package request

import (
	"app/chapter-struct-mapper-pool/pool"
)

var poolArticleGet = pool.New(func() *ArticleGetWithPool {
	return new(ArticleGetWithPool)
})
var poolArticleAll = pool.New(func() *ArticleAllWithPool {
	return new(ArticleAllWithPool)
})
