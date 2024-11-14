package repository

import (
	"context"
	"elegantGo/orm-crud-2/repository/ent"
	"elegantGo/orm-crud-2/repository/ent/post"
)

type Post struct {
	repo
}

func NewPost(db *ent.Client) *Post {
	repo := new(Post)
	repo.db = db
	return repo
}

func (repo *Post) Fetch(ctx context.Context, id int) *ent.Post {
	return repo.fetchOne(ctx, repo.db, func(opt *ent.PostQuery) {
		opt.Where(post.ID(id))
	})
}
func (repo *Post) FetchOne(ctx context.Context, optionFunc func(opt *ent.PostQuery)) *ent.Post {
	return repo.fetchOne(ctx, repo.db, optionFunc)
}

func (repo *Post) fetchOne(ctx context.Context, db *ent.Client, optionFunc func(opt *ent.PostQuery)) *ent.Post {
	builder := db.Post.Query()
	optionFunc(builder)
	return builder.FirstX(ctx)
}

func (repo *Post) FetchMany(ctx context.Context, optionFunc func(opt *ent.PostQuery)) ent.Posts {
	return repo.fetchMany(ctx, repo.db, optionFunc)
}

func (repo *Post) fetchMany(ctx context.Context, db *ent.Client, optionFunc func(opt *ent.PostQuery)) ent.Posts {
	builder := db.Post.Query()
	optionFunc(builder)
	return builder.AllX(ctx)
}
