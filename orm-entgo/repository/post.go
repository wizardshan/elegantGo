package repository

import (
	"context"
	"elegantGo/orm-entgo/repository/ent"
	"elegantGo/orm-entgo/repository/ent/comment"
	"elegantGo/orm-entgo/repository/ent/post"
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
	return repo.db.Post.Query().WithUser().WithComments(func(opts *ent.CommentQuery) {
		opts.WithUser()
	}).Where(post.ID(id)).FirstX(ctx)
}

func (repo *Post) FetchMany(ctx context.Context) []*ent.Post {
	return repo.db.Debug().Post.Query().WithUser().Order(ent.Desc(post.FieldCreateTime)).AllX(ctx)
}

func (repo *Post) Comments(ctx context.Context) []*ent.Comment {
	return repo.db.Comment.Query().WithUser().WithPost().Order(ent.Desc(comment.FieldCreateTime)).AllX(ctx)
}
