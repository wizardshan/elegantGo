package request

import (
	"app/chapter5.0/pool"
)

var poolArticleGet = pool.New(func() *ArticleGetWithPool {
	return new(ArticleGetWithPool)
})
var poolArticleAll = pool.New(func() *ArticleAllWithPool {
	return new(ArticleAllWithPool)
})
