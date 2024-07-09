package request

import "elegantGo/chapter-struct-mapper-pool/pkg/pool"

var poolArticle = pool.New(func() *ArticlePool {
	return new(ArticlePool)
})
var poolArticles = pool.New(func() *ArticlesPool {
	return new(ArticlesPool)
})

type ArticlePool struct {
	ID int `form:"id"`
}

func NewArticle() *ArticlePool {
	//fmt.Println("req poolArticle.Get")
	return poolArticle.Get()
}

func (reqArticle *ArticlePool) Recycle() {
	//fmt.Println("req poolArticle.Put")
	poolArticle.Put(reqArticle)
}

type ArticlesPool struct {
}

func NewArticles() *ArticlesPool {
	//fmt.Println("req poolArticleAll.Get")
	return poolArticles.Get()
}

func (reqArticle *ArticlesPool) Recycle() {
	//fmt.Println("req poolArticleAll.Put")
	poolArticles.Put(reqArticle)
}
