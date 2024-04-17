package request

import (
	"app/chapter1.4/controller/request/article"
	"app/chapter1.4/controller/request/column"
)

type ColumnCreateArticle struct {
	IDField
	article.TitleField
}

type ColumnCreate struct {
	column.TitleField
	Article  *ColumnCreateArticle  `form:"article"`
	Articles []ColumnCreateArticle `form:"articles"`
}

type ColumnDetail struct {
	IDField
	ArticleIsImportant *bool `form:"articleIsImportant"`
}
