package request

import (
	"app/chapter-param-validator-5/controller/request/article"
	"app/chapter-param-validator-5/controller/request/column"
)

type ColumnCreateArticle struct {
	IDField
	article.TitleField
}

type ColumnCreate struct {
	column.TitleField
	ArticleOnTop *ColumnCreateArticle
	Articles     []ColumnCreateArticle
}

type ColumnOne struct {
	IDField
	Important *bool
}
