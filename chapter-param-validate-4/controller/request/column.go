package request

import (
	"elegantGo/chapter-param-validate-4/controller/request/article"
	"elegantGo/chapter-param-validate-4/controller/request/column"
)

type ColumnCreateArticle struct {
	IDField
	article.TitleField
}

type ColumnCreate struct {
	column.TitleField
	Article  *ColumnCreateArticle
	Articles []*ColumnCreateArticle
}

type ColumnOne struct {
	IDField
	Important *bool
}
