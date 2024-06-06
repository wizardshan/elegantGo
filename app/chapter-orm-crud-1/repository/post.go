package repository

import (
	"app/chapter-orm-crud-1/repository/ent"
	"app/chapter-orm-crud-1/repository/ent/post"
	"context"
)

type Post struct {
	db *ent.Client
}

func NewPost(db *ent.Client) *Post {
	repo := new(Post)
	repo.db = db
	return repo
}

func (repo *Post) Fetch(ctx context.Context, id int) *ent.Post {
	return repo.fetchOne(ctx, repo.db, func(builder *ent.PostQuery) {
		builder.Where(post.ID(id))
	})
}
func (repo *Post) FetchOne(ctx context.Context, optionFunc func(builder *ent.PostQuery)) *ent.Post {
	return repo.fetchOne(ctx, repo.db, optionFunc)
}

func (repo *Post) fetchOne(ctx context.Context, db *ent.Client, optionFunc func(builder *ent.PostQuery)) *ent.Post {
	builder := db.Post.Query()
	optionFunc(builder)
	return builder.FirstX(ctx)
}

func (repo *Post) FetchMany(ctx context.Context, optionFunc func(builder *ent.PostQuery)) ent.Posts {
	return repo.fetchMany(ctx, repo.db, optionFunc)
}

func (repo *Post) fetchMany(ctx context.Context, db *ent.Client, optionFunc func(builder *ent.PostQuery)) ent.Posts {
	builder := db.Post.Query()
	optionFunc(builder)
	return builder.AllX(ctx)
}
