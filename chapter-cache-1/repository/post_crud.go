package repository

import (
	"context"
	"elegantGo/chapter-cache-1/domain"
	"elegantGo/chapter-cache-1/repository/ent"
	"elegantGo/chapter-cache-1/repository/ent/post"
)

type Post struct {
	repo
}

func NewPost(db *ent.Client) *Post {
	repo := new(Post)
	repo.db = db
	return repo
}

func (repo *Post) Fetch(ctx context.Context, id int) *domain.Post {
	return repo.fetchOne(ctx, repo.db, func(builder *ent.PostQuery) {
		builder.Where(post.ID(id))
	}).Mapper()
}
func (repo *Post) FetchOne(ctx context.Context, optionFunc func(builder *ent.PostQuery)) *domain.Post {
	return repo.fetchOne(ctx, repo.db, optionFunc).Mapper()
}

func (repo *Post) fetchOne(ctx context.Context, db *ent.Client, optionFunc func(builder *ent.PostQuery)) *ent.Post {
	builder := db.Post.Query()
	optionFunc(builder)
	return builder.FirstX(ctx)
}

func (repo *Post) FetchMany(ctx context.Context, optionFunc func(builder *ent.PostQuery)) domain.Posts {
	return repo.fetchMany(ctx, repo.db, optionFunc).Mapper()
}

func (repo *Post) fetchMany(ctx context.Context, db *ent.Client, optionFunc func(builder *ent.PostQuery)) ent.Posts {
	builder := db.Post.Query()
	optionFunc(builder)
	return builder.AllX(ctx)
}

func (repo *Post) Count(ctx context.Context, optionFunc func(builder *ent.PostQuery)) int {
	return repo.count(ctx, repo.db, optionFunc)
}

func (repo *Post) count(ctx context.Context, db *ent.Client, optionFunc func(builder *ent.PostQuery)) int {
	builder := db.Post.Query()
	optionFunc(builder)
	return builder.CountX(ctx)
}

func (repo *Post) Exist(ctx context.Context, optionFunc func(builder *ent.PostQuery)) bool {
	return repo.exist(ctx, repo.db, optionFunc)
}

func (repo *Post) exist(ctx context.Context, db *ent.Client, optionFunc func(builder *ent.PostQuery)) bool {
	builder := db.Post.Query()
	optionFunc(builder)
	return builder.ExistX(ctx)
}

func (repo *Post) Create(ctx context.Context, optionFunc func(builder *ent.PostCreate)) *domain.Post {
	return repo.create(ctx, repo.db, optionFunc).Mapper()
}

func (repo *Post) create(ctx context.Context, db *ent.Client, optionFunc func(builder *ent.PostCreate)) *ent.Post {
	builder := db.Post.Create()
	optionFunc(builder)
	return builder.SaveX(ctx)
}

func (repo *Post) Update(ctx context.Context, optionFunc func(builder *ent.PostUpdate)) int {
	return repo.update(ctx, repo.db, optionFunc)
}

func (repo *Post) update(ctx context.Context, db *ent.Client, optionFunc func(builder *ent.PostUpdate)) int {
	builder := db.Post.Update()
	optionFunc(builder)
	return builder.SaveX(ctx)
}

func (repo *Post) Delete(ctx context.Context, optionFunc func(builder *ent.PostDelete)) int {
	return repo.delete(ctx, repo.db, optionFunc)
}

func (repo *Post) delete(ctx context.Context, db *ent.Client, optionFunc func(builder *ent.PostDelete)) int {
	builder := db.Post.Delete()
	optionFunc(builder)
	return builder.ExecX(ctx)
}
