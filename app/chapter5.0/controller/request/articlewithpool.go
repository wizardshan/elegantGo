package request

type ArticleGetWithPool struct {
	ID int `form:"id"`
}

func NewArticleGetWithPool() *ArticleGetWithPool {
	//fmt.Println("req poolArticle.Get")
	return poolArticleGet.Get()
}

func (reqArticle *ArticleGetWithPool) Put() {
	//fmt.Println("req poolArticle.Put")
	poolArticleGet.Put(reqArticle)
}

type ArticleAllWithPool struct {
}

func NewArticleAllWithPool() *ArticleAllWithPool {
	//fmt.Println("req poolArticleAll.Get")
	return poolArticleAll.Get()
}

func (reqArticle *ArticleAllWithPool) Put() {
	//fmt.Println("req poolArticleAll.Put")
	poolArticleAll.Put(reqArticle)
}
