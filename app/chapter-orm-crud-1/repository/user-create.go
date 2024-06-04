package repository

import (
	"app/chapter-orm-crud-1/repository/ent"
	"context"
)

func (repo *User) Create(ctx context.Context, optionFunc func(builder *ent.UserCreate)) *ent.User {
	builder := repo.db.User.Create()
	optionFunc(builder)
	return builder.SaveX(ctx)
}
