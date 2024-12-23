package repository

import (
	"context"
	"elegantGo/orm-crud-2/repository/ent"
	"elegantGo/orm-crud-2/repository/ent/user"
)

func (repo *User) Fetch(ctx context.Context, id int) *ent.User {
	return repo.fetchOne(ctx, repo.db, func(opt *ent.UserQuery) {
		opt.Where(user.ID(id))
	})
}
func (repo *User) FetchOne(ctx context.Context, optionFunc func(opt *ent.UserQuery)) *ent.User {
	return repo.fetchOne(ctx, repo.db, optionFunc)
}

func (repo *User) fetchOne(ctx context.Context, db *ent.Client, optionFunc func(opt *ent.UserQuery)) *ent.User {
	builder := db.User.Query()
	optionFunc(builder)
	return builder.FirstX(ctx)
}

func (repo *User) FetchMany(ctx context.Context, optionFunc func(opt *ent.UserQuery)) ent.Users {
	return repo.fetchMany(ctx, repo.db, optionFunc)
}

func (repo *User) fetchMany(ctx context.Context, db *ent.Client, optionFunc func(opt *ent.UserQuery)) ent.Users {
	builder := db.User.Query()
	optionFunc(builder)
	return builder.AllX(ctx)
}

func (repo *User) Count(ctx context.Context, optionFunc func(opt *ent.UserQuery)) int {
	return repo.count(ctx, repo.db, optionFunc)
}

func (repo *User) count(ctx context.Context, db *ent.Client, optionFunc func(opt *ent.UserQuery)) int {
	builder := db.User.Query()
	optionFunc(builder)
	return builder.CountX(ctx)
}

func (repo *User) Exist(ctx context.Context, optionFunc func(builder *ent.UserQuery)) bool {
	return repo.exist(ctx, repo.db, optionFunc)
}

func (repo *User) exist(ctx context.Context, db *ent.Client, optionFunc func(builder *ent.UserQuery)) bool {
	builder := db.User.Query()
	optionFunc(builder)
	return builder.ExistX(ctx)
}
