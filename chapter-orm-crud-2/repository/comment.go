package repository

import (
	"context"
	"elegantGo/chapter-orm-crud-2/repository/ent"
	"elegantGo/chapter-orm-crud-2/repository/ent/comment"
)

type Comment struct {
	repo
}

func NewComment(db *ent.Client) *Comment {
	repo := new(Comment)
	repo.db = db
	return repo
}

func (repo *Comment) Fetch(ctx context.Context, id int) *ent.Comment {
	return repo.fetchOne(ctx, repo.db, func(opt *ent.CommentQuery) {
		opt.Where(comment.ID(id))
	})
}
func (repo *Comment) FetchOne(ctx context.Context, optionFunc func(opt *ent.CommentQuery)) *ent.Comment {
	return repo.fetchOne(ctx, repo.db, optionFunc)
}

func (repo *Comment) fetchOne(ctx context.Context, db *ent.Client, optionFunc func(opt *ent.CommentQuery)) *ent.Comment {
	builder := db.Comment.Query()
	optionFunc(builder)
	return builder.FirstX(ctx)
}

func (repo *Comment) FetchMany(ctx context.Context, optionFunc func(opt *ent.CommentQuery)) ent.Comments {
	return repo.fetchMany(ctx, repo.db, optionFunc)
}

func (repo *Comment) fetchMany(ctx context.Context, db *ent.Client, optionFunc func(opt *ent.CommentQuery)) ent.Comments {
	builder := db.Comment.Query()
	optionFunc(builder)
	return builder.AllX(ctx)
}
