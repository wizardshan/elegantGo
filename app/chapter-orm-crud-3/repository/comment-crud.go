package repository

import (
	"app/chapter-orm-crud-3/repository/ent"
	"app/chapter-orm-crud-3/repository/ent/comment"
	"context"
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
	return repo.fetchOne(ctx, repo.db, func(builder *ent.CommentQuery) {
		builder.Where(comment.ID(id))
	})
}
func (repo *Comment) FetchOne(ctx context.Context, optionFunc func(builder *ent.CommentQuery)) *ent.Comment {
	return repo.fetchOne(ctx, repo.db, optionFunc)
}

func (repo *Comment) fetchOne(ctx context.Context, db *ent.Client, optionFunc func(builder *ent.CommentQuery)) *ent.Comment {
	builder := db.Comment.Query()
	optionFunc(builder)
	return builder.FirstX(ctx)
}

func (repo *Comment) FetchMany(ctx context.Context, optionFunc func(builder *ent.CommentQuery)) ent.Comments {
	return repo.fetchMany(ctx, repo.db, optionFunc)
}

func (repo *Comment) fetchMany(ctx context.Context, db *ent.Client, optionFunc func(builder *ent.CommentQuery)) ent.Comments {
	builder := db.Comment.Query()
	optionFunc(builder)
	return builder.AllX(ctx)
}

func (repo *Comment) Count(ctx context.Context, optionFunc func(builder *ent.CommentQuery)) int {
	return repo.count(ctx, repo.db, optionFunc)
}

func (repo *Comment) count(ctx context.Context, db *ent.Client, optionFunc func(builder *ent.CommentQuery)) int {
	builder := db.Comment.Query()
	optionFunc(builder)
	return builder.CountX(ctx)
}

func (repo *Comment) Exist(ctx context.Context, optionFunc func(builder *ent.CommentQuery)) bool {
	return repo.exist(ctx, repo.db, optionFunc)
}

func (repo *Comment) exist(ctx context.Context, db *ent.Client, optionFunc func(builder *ent.CommentQuery)) bool {
	builder := db.Comment.Query()
	optionFunc(builder)
	return builder.ExistX(ctx)
}

func (repo *Comment) Create(ctx context.Context, optionFunc func(builder *ent.CommentCreate)) *ent.Comment {
	return repo.create(ctx, repo.db, optionFunc)
}

func (repo *Comment) create(ctx context.Context, db *ent.Client, optionFunc func(builder *ent.CommentCreate)) *ent.Comment {
	builder := db.Comment.Create()
	optionFunc(builder)
	return builder.SaveX(ctx)
}

func (repo *Comment) Update(ctx context.Context, optionFunc func(builder *ent.CommentUpdate)) int {
	return repo.update(ctx, repo.db, optionFunc)
}

func (repo *Comment) update(ctx context.Context, db *ent.Client, optionFunc func(builder *ent.CommentUpdate)) int {
	builder := db.Comment.Update()
	optionFunc(builder)
	return builder.SaveX(ctx)
}

func (repo *Comment) Delete(ctx context.Context, optionFunc func(builder *ent.CommentDelete)) int {
	return repo.delete(ctx, repo.db, optionFunc)
}

func (repo *Comment) delete(ctx context.Context, db *ent.Client, optionFunc func(builder *ent.CommentDelete)) int {
	builder := db.Comment.Delete()
	optionFunc(builder)
	return builder.ExecX(ctx)
}