package repository

import (
	"context"
	"elegantGo/orm-crud-1/repository/ent"
)

func (repo *User) Delete(ctx context.Context, optionFunc func(opt *ent.UserDelete)) int {
	builder := repo.db.User.Delete()
	optionFunc(builder)
	return builder.ExecX(ctx)
}
