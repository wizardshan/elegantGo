package repository

import (
	"context"
	"elegantGo/chapter-orm-crud-2/repository/ent"
)

func (repo *User) Create(ctx context.Context, optionFunc func(builder *ent.UserCreate)) *ent.User {
	return repo.create(ctx, repo.db, optionFunc)
}

func (repo *User) create(ctx context.Context, db *ent.Client, optionFunc func(builder *ent.UserCreate)) *ent.User {
	builder := db.User.Create()
	optionFunc(builder)
	return builder.SaveX(ctx)
}
