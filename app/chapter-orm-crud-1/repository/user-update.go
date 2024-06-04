package repository

import (
	"app/chapter-orm-crud-1/repository/ent"
	"context"
)

func (repo *User) Update(ctx context.Context, optionFunc func(builder *ent.UserUpdate)) int {
	builder := repo.db.User.Update()
	optionFunc(builder)
	return builder.SaveX(ctx)
}
