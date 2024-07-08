package repository

import (
	"app/chapter-param-validator-sql-injection/repository/entity"
	"context"
	"database/sql"
	"fmt"
)

type Article struct {
	db *sql.DB
}

func NewArticle(db *sql.DB) *Article {
	repo := new(Article)
	repo.db = db
	return repo
}

func (repo *Article) Get(ctx context.Context, hashID string) *entity.Article {
	query := fmt.Sprintf("select hash_id, title, content from articles where hash_id='%s'", hashID)
	row := repo.db.QueryRowContext(ctx, query)

	//query := "select hash_id, title, content from articles where hash_id=?"
	//row := repo.db.QueryRowContext(ctx, query, hashID)

	var article entity.Article
	article.HashIDQuery = hashID
	article.SQL = query
	if row.Err() != nil {
		article.Err = row.Err().Error()
	}
	row.Scan(&article.HashID, &article.Title, &article.Content)
	return &article
}
