package repository

import (
	"context"
	"elegantGo/chapter-orm-crud-1/repository/ent"
)

func (repo *User) Update(ctx context.Context, optionFunc func(opt *ent.UserUpdate)) int {
	builder := repo.db.User.Update()
	optionFunc(builder)
	return builder.SaveX(ctx)
}
