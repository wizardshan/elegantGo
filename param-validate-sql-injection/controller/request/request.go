package request

type HashIDField struct {
	HashID string `binding:"required"`
}

type KeywordField struct {
	Keyword string `binding:"required,sqlinject"`
}
