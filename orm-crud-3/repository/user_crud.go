package repository

import (
	"context"
	"elegantGo/orm-crud-3/repository/ent"
	"elegantGo/orm-crud-3/repository/ent/user"
)

type User struct {
	repo
}

func NewUser(db *ent.Client) *User {
	repo := new(User)
	repo.db = db
	return repo
}

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

func (repo *User) Exist(ctx context.Context, optionFunc func(opt *ent.UserQuery)) bool {
	return repo.exist(ctx, repo.db, optionFunc)
}

func (repo *User) exist(ctx context.Context, db *ent.Client, optionFunc func(opt *ent.UserQuery)) bool {
	builder := db.User.Query()
	optionFunc(builder)
	return builder.ExistX(ctx)
}

func (repo *User) Create(ctx context.Context, optionFunc func(opt *ent.UserCreate)) *ent.User {
	return repo.create(ctx, repo.db, optionFunc)
}

func (repo *User) create(ctx context.Context, db *ent.Client, optionFunc func(opt *ent.UserCreate)) *ent.User {
	builder := db.User.Create()
	optionFunc(builder)
	return builder.SaveX(ctx)
}

func (repo *User) Update(ctx context.Context, optionFunc func(opt *ent.UserUpdate)) int {
	return repo.update(ctx, repo.db, optionFunc)
}

func (repo *User) update(ctx context.Context, db *ent.Client, optionFunc func(opt *ent.UserUpdate)) int {
	builder := db.User.Update()
	optionFunc(builder)
	return builder.SaveX(ctx)
}

func (repo *User) Delete(ctx context.Context, optionFunc func(opt *ent.UserDelete)) int {
	return repo.delete(ctx, repo.db, optionFunc)
}

func (repo *User) delete(ctx context.Context, db *ent.Client, optionFunc func(opt *ent.UserDelete)) int {
	builder := db.User.Delete()
	optionFunc(builder)
	return builder.ExecX(ctx)
}
