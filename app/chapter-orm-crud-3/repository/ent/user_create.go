// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/chapter-orm-crud-3/repository/ent/comment"
	"app/chapter-orm-crud-3/repository/ent/post"
	"app/chapter-orm-crud-3/repository/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (uc *UserCreate) SetCreateTime(t time.Time) *UserCreate {
	uc.mutation.SetCreateTime(t)
	return uc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreateTime(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreateTime(*t)
	}
	return uc
}

// SetUpdateTime sets the "update_time" field.
func (uc *UserCreate) SetUpdateTime(t time.Time) *UserCreate {
	uc.mutation.SetUpdateTime(t)
	return uc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdateTime(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdateTime(*t)
	}
	return uc
}

// SetHashID sets the "hash_id" field.
func (uc *UserCreate) SetHashID(s string) *UserCreate {
	uc.mutation.SetHashID(s)
	return uc
}

// SetNillableHashID sets the "hash_id" field if the given value is not nil.
func (uc *UserCreate) SetNillableHashID(s *string) *UserCreate {
	if s != nil {
		uc.SetHashID(*s)
	}
	return uc
}

// SetMobile sets the "mobile" field.
func (uc *UserCreate) SetMobile(s string) *UserCreate {
	uc.mutation.SetMobile(s)
	return uc
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (uc *UserCreate) SetNillableMobile(s *string) *UserCreate {
	if s != nil {
		uc.SetMobile(*s)
	}
	return uc
}

// SetPassword sets the "password" field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uc *UserCreate) SetNillablePassword(s *string) *UserCreate {
	if s != nil {
		uc.SetPassword(*s)
	}
	return uc
}

// SetLevel sets the "level" field.
func (uc *UserCreate) SetLevel(i int) *UserCreate {
	uc.mutation.SetLevel(i)
	return uc
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (uc *UserCreate) SetNillableLevel(i *int) *UserCreate {
	if i != nil {
		uc.SetLevel(*i)
	}
	return uc
}

// SetNickname sets the "nickname" field.
func (uc *UserCreate) SetNickname(s string) *UserCreate {
	uc.mutation.SetNickname(s)
	return uc
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (uc *UserCreate) SetNillableNickname(s *string) *UserCreate {
	if s != nil {
		uc.SetNickname(*s)
	}
	return uc
}

// SetAvatar sets the "avatar" field.
func (uc *UserCreate) SetAvatar(s string) *UserCreate {
	uc.mutation.SetAvatar(s)
	return uc
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uc *UserCreate) SetNillableAvatar(s *string) *UserCreate {
	if s != nil {
		uc.SetAvatar(*s)
	}
	return uc
}

// SetBio sets the "bio" field.
func (uc *UserCreate) SetBio(s string) *UserCreate {
	uc.mutation.SetBio(s)
	return uc
}

// SetNillableBio sets the "bio" field if the given value is not nil.
func (uc *UserCreate) SetNillableBio(s *string) *UserCreate {
	if s != nil {
		uc.SetBio(*s)
	}
	return uc
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (uc *UserCreate) AddCommentIDs(ids ...int) *UserCreate {
	uc.mutation.AddCommentIDs(ids...)
	return uc
}

// AddComments adds the "comments" edges to the Comment entity.
func (uc *UserCreate) AddComments(c ...*Comment) *UserCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddCommentIDs(ids...)
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (uc *UserCreate) AddPostIDs(ids ...int) *UserCreate {
	uc.mutation.AddPostIDs(ids...)
	return uc
}

// AddPosts adds the "posts" edges to the Post entity.
func (uc *UserCreate) AddPosts(p ...*Post) *UserCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uc.AddPostIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.CreateTime(); !ok {
		v := user.DefaultCreateTime()
		uc.mutation.SetCreateTime(v)
	}
	if _, ok := uc.mutation.UpdateTime(); !ok {
		v := user.DefaultUpdateTime()
		uc.mutation.SetUpdateTime(v)
	}
	if _, ok := uc.mutation.HashID(); !ok {
		v := user.DefaultHashID
		uc.mutation.SetHashID(v)
	}
	if _, ok := uc.mutation.Mobile(); !ok {
		v := user.DefaultMobile
		uc.mutation.SetMobile(v)
	}
	if _, ok := uc.mutation.Password(); !ok {
		v := user.DefaultPassword
		uc.mutation.SetPassword(v)
	}
	if _, ok := uc.mutation.Level(); !ok {
		v := user.DefaultLevel
		uc.mutation.SetLevel(v)
	}
	if _, ok := uc.mutation.Nickname(); !ok {
		v := user.DefaultNickname
		uc.mutation.SetNickname(v)
	}
	if _, ok := uc.mutation.Avatar(); !ok {
		v := user.DefaultAvatar
		uc.mutation.SetAvatar(v)
	}
	if _, ok := uc.mutation.Bio(); !ok {
		v := user.DefaultBio
		uc.mutation.SetBio(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "User.create_time"`)}
	}
	if _, ok := uc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "User.update_time"`)}
	}
	if _, ok := uc.mutation.HashID(); !ok {
		return &ValidationError{Name: "hash_id", err: errors.New(`ent: missing required field "User.hash_id"`)}
	}
	if _, ok := uc.mutation.Mobile(); !ok {
		return &ValidationError{Name: "mobile", err: errors.New(`ent: missing required field "User.mobile"`)}
	}
	if _, ok := uc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "User.password"`)}
	}
	if _, ok := uc.mutation.Level(); !ok {
		return &ValidationError{Name: "level", err: errors.New(`ent: missing required field "User.level"`)}
	}
	if _, ok := uc.mutation.Nickname(); !ok {
		return &ValidationError{Name: "nickname", err: errors.New(`ent: missing required field "User.nickname"`)}
	}
	if _, ok := uc.mutation.Avatar(); !ok {
		return &ValidationError{Name: "avatar", err: errors.New(`ent: missing required field "User.avatar"`)}
	}
	if _, ok := uc.mutation.Bio(); !ok {
		return &ValidationError{Name: "bio", err: errors.New(`ent: missing required field "User.bio"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	)
	_spec.OnConflict = uc.conflict
	if value, ok := uc.mutation.CreateTime(); ok {
		_spec.SetField(user.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := uc.mutation.UpdateTime(); ok {
		_spec.SetField(user.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := uc.mutation.HashID(); ok {
		_spec.SetField(user.FieldHashID, field.TypeString, value)
		_node.HashID = value
	}
	if value, ok := uc.mutation.Mobile(); ok {
		_spec.SetField(user.FieldMobile, field.TypeString, value)
		_node.Mobile = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := uc.mutation.Level(); ok {
		_spec.SetField(user.FieldLevel, field.TypeInt, value)
		_node.Level = value
	}
	if value, ok := uc.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
		_node.Nickname = value
	}
	if value, ok := uc.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
		_node.Avatar = value
	}
	if value, ok := uc.mutation.Bio(); ok {
		_spec.SetField(user.FieldBio, field.TypeString, value)
		_node.Bio = value
	}
	if nodes := uc.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CommentsTable,
			Columns: []string{user.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PostsTable,
			Columns: []string{user.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (uc *UserCreate) OnConflict(opts ...sql.ConflictOption) *UserUpsertOne {
	uc.conflict = opts
	return &UserUpsertOne{
		create: uc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (uc *UserCreate) OnConflictColumns(columns ...string) *UserUpsertOne {
	uc.conflict = append(uc.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertOne{
		create: uc,
	}
}

type (
	// UserUpsertOne is the builder for "upsert"-ing
	//  one User node.
	UserUpsertOne struct {
		create *UserCreate
	}

	// UserUpsert is the "OnConflict" setter.
	UserUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *UserUpsert) SetUpdateTime(v time.Time) *UserUpsert {
	u.Set(user.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserUpsert) UpdateUpdateTime() *UserUpsert {
	u.SetExcluded(user.FieldUpdateTime)
	return u
}

// SetHashID sets the "hash_id" field.
func (u *UserUpsert) SetHashID(v string) *UserUpsert {
	u.Set(user.FieldHashID, v)
	return u
}

// UpdateHashID sets the "hash_id" field to the value that was provided on create.
func (u *UserUpsert) UpdateHashID() *UserUpsert {
	u.SetExcluded(user.FieldHashID)
	return u
}

// SetMobile sets the "mobile" field.
func (u *UserUpsert) SetMobile(v string) *UserUpsert {
	u.Set(user.FieldMobile, v)
	return u
}

// UpdateMobile sets the "mobile" field to the value that was provided on create.
func (u *UserUpsert) UpdateMobile() *UserUpsert {
	u.SetExcluded(user.FieldMobile)
	return u
}

// SetPassword sets the "password" field.
func (u *UserUpsert) SetPassword(v string) *UserUpsert {
	u.Set(user.FieldPassword, v)
	return u
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *UserUpsert) UpdatePassword() *UserUpsert {
	u.SetExcluded(user.FieldPassword)
	return u
}

// SetLevel sets the "level" field.
func (u *UserUpsert) SetLevel(v int) *UserUpsert {
	u.Set(user.FieldLevel, v)
	return u
}

// UpdateLevel sets the "level" field to the value that was provided on create.
func (u *UserUpsert) UpdateLevel() *UserUpsert {
	u.SetExcluded(user.FieldLevel)
	return u
}

// AddLevel adds v to the "level" field.
func (u *UserUpsert) AddLevel(v int) *UserUpsert {
	u.Add(user.FieldLevel, v)
	return u
}

// SetNickname sets the "nickname" field.
func (u *UserUpsert) SetNickname(v string) *UserUpsert {
	u.Set(user.FieldNickname, v)
	return u
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *UserUpsert) UpdateNickname() *UserUpsert {
	u.SetExcluded(user.FieldNickname)
	return u
}

// SetAvatar sets the "avatar" field.
func (u *UserUpsert) SetAvatar(v string) *UserUpsert {
	u.Set(user.FieldAvatar, v)
	return u
}

// UpdateAvatar sets the "avatar" field to the value that was provided on create.
func (u *UserUpsert) UpdateAvatar() *UserUpsert {
	u.SetExcluded(user.FieldAvatar)
	return u
}

// SetBio sets the "bio" field.
func (u *UserUpsert) SetBio(v string) *UserUpsert {
	u.Set(user.FieldBio, v)
	return u
}

// UpdateBio sets the "bio" field to the value that was provided on create.
func (u *UserUpsert) UpdateBio() *UserUpsert {
	u.SetExcluded(user.FieldBio)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserUpsertOne) UpdateNewValues() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(user.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserUpsertOne) Ignore() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertOne) DoNothing() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreate.OnConflict
// documentation for more info.
func (u *UserUpsertOne) Update(set func(*UserUpsert)) *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *UserUpsertOne) SetUpdateTime(v time.Time) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateUpdateTime() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetHashID sets the "hash_id" field.
func (u *UserUpsertOne) SetHashID(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetHashID(v)
	})
}

// UpdateHashID sets the "hash_id" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateHashID() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateHashID()
	})
}

// SetMobile sets the "mobile" field.
func (u *UserUpsertOne) SetMobile(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetMobile(v)
	})
}

// UpdateMobile sets the "mobile" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateMobile() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateMobile()
	})
}

// SetPassword sets the "password" field.
func (u *UserUpsertOne) SetPassword(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *UserUpsertOne) UpdatePassword() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdatePassword()
	})
}

// SetLevel sets the "level" field.
func (u *UserUpsertOne) SetLevel(v int) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetLevel(v)
	})
}

// AddLevel adds v to the "level" field.
func (u *UserUpsertOne) AddLevel(v int) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.AddLevel(v)
	})
}

// UpdateLevel sets the "level" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateLevel() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateLevel()
	})
}

// SetNickname sets the "nickname" field.
func (u *UserUpsertOne) SetNickname(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetNickname(v)
	})
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateNickname() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateNickname()
	})
}

// SetAvatar sets the "avatar" field.
func (u *UserUpsertOne) SetAvatar(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetAvatar(v)
	})
}

// UpdateAvatar sets the "avatar" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateAvatar() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateAvatar()
	})
}

// SetBio sets the "bio" field.
func (u *UserUpsertOne) SetBio(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetBio(v)
	})
}

// UpdateBio sets the "bio" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateBio() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateBio()
	})
}

// Exec executes the query.
func (u *UserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
	conflict []sql.ConflictOption
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserUpsertBulk {
	ucb.conflict = opts
	return &UserUpsertBulk{
		create: ucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflictColumns(columns ...string) *UserUpsertBulk {
	ucb.conflict = append(ucb.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertBulk{
		create: ucb,
	}
}

// UserUpsertBulk is the builder for "upsert"-ing
// a bulk of User nodes.
type UserUpsertBulk struct {
	create *UserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserUpsertBulk) UpdateNewValues() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(user.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserUpsertBulk) Ignore() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertBulk) DoNothing() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreateBulk.OnConflict
// documentation for more info.
func (u *UserUpsertBulk) Update(set func(*UserUpsert)) *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *UserUpsertBulk) SetUpdateTime(v time.Time) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateUpdateTime() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetHashID sets the "hash_id" field.
func (u *UserUpsertBulk) SetHashID(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetHashID(v)
	})
}

// UpdateHashID sets the "hash_id" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateHashID() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateHashID()
	})
}

// SetMobile sets the "mobile" field.
func (u *UserUpsertBulk) SetMobile(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetMobile(v)
	})
}

// UpdateMobile sets the "mobile" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateMobile() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateMobile()
	})
}

// SetPassword sets the "password" field.
func (u *UserUpsertBulk) SetPassword(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdatePassword() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdatePassword()
	})
}

// SetLevel sets the "level" field.
func (u *UserUpsertBulk) SetLevel(v int) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetLevel(v)
	})
}

// AddLevel adds v to the "level" field.
func (u *UserUpsertBulk) AddLevel(v int) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.AddLevel(v)
	})
}

// UpdateLevel sets the "level" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateLevel() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateLevel()
	})
}

// SetNickname sets the "nickname" field.
func (u *UserUpsertBulk) SetNickname(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetNickname(v)
	})
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateNickname() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateNickname()
	})
}

// SetAvatar sets the "avatar" field.
func (u *UserUpsertBulk) SetAvatar(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetAvatar(v)
	})
}

// UpdateAvatar sets the "avatar" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateAvatar() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateAvatar()
	})
}

// SetBio sets the "bio" field.
func (u *UserUpsertBulk) SetBio(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetBio(v)
	})
}

// UpdateBio sets the "bio" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateBio() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateBio()
	})
}

// Exec executes the query.
func (u *UserUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
