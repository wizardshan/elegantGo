package request

type ArticlePool struct {
	ID int `form:"id"`
}

func NewArticlePool() *ArticlePool {
	//fmt.Println("req poolArticle.Get")
	return poolArticle.Get()
}

func (reqArticle *ArticlePool) Put() {
	//fmt.Println("req poolArticle.Put")
	poolArticle.Put(reqArticle)
}

type ArticlesPool struct {
}

func NewArticlesPool() *ArticlesPool {
	//fmt.Println("req poolArticleAll.Get")
	return poolArticles.Get()
}

func (reqArticle *ArticlesPool) Put() {
	//fmt.Println("req poolArticleAll.Put")
	poolArticles.Put(reqArticle)
}
