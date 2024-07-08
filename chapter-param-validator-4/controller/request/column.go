package request

import (
	"elegantGo/chapter-param-validator-4/controller/request/article"
	"elegantGo/chapter-param-validator-4/controller/request/column"
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
