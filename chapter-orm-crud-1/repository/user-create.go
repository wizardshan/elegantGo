package repository

import (
	"context"
	"elegantGo/chapter-orm-crud-1/repository/ent"
)

func (repo *User) Create(ctx context.Context, optionFunc func(builder *ent.UserCreate)) *ent.User {
	builder := repo.db.User.Create()
	optionFunc(builder)
	return builder.SaveX(ctx)
}
