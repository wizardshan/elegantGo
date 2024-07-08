package request

import (
	"elegantGo/chapter-struct-mapper-pool/pool"
)

var poolArticle = pool.New(func() *ArticlePool {
	return new(ArticlePool)
})
var poolArticles = pool.New(func() *ArticlesPool {
	return new(ArticlesPool)
})
