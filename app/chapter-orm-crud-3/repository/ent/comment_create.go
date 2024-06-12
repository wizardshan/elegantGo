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

// CommentCreate is the builder for creating a Comment entity.
type CommentCreate struct {
	config
	mutation *CommentMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (cc *CommentCreate) SetCreateTime(t time.Time) *CommentCreate {
	cc.mutation.SetCreateTime(t)
	return cc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (cc *CommentCreate) SetNillableCreateTime(t *time.Time) *CommentCreate {
	if t != nil {
		cc.SetCreateTime(*t)
	}
	return cc
}

// SetUpdateTime sets the "update_time" field.
func (cc *CommentCreate) SetUpdateTime(t time.Time) *CommentCreate {
	cc.mutation.SetUpdateTime(t)
	return cc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (cc *CommentCreate) SetNillableUpdateTime(t *time.Time) *CommentCreate {
	if t != nil {
		cc.SetUpdateTime(*t)
	}
	return cc
}

// SetUserID sets the "user_id" field.
func (cc *CommentCreate) SetUserID(i int) *CommentCreate {
	cc.mutation.SetUserID(i)
	return cc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cc *CommentCreate) SetNillableUserID(i *int) *CommentCreate {
	if i != nil {
		cc.SetUserID(*i)
	}
	return cc
}

// SetPostID sets the "post_id" field.
func (cc *CommentCreate) SetPostID(i int) *CommentCreate {
	cc.mutation.SetPostID(i)
	return cc
}

// SetNillablePostID sets the "post_id" field if the given value is not nil.
func (cc *CommentCreate) SetNillablePostID(i *int) *CommentCreate {
	if i != nil {
		cc.SetPostID(*i)
	}
	return cc
}

// SetContent sets the "content" field.
func (cc *CommentCreate) SetContent(s string) *CommentCreate {
	cc.mutation.SetContent(s)
	return cc
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (cc *CommentCreate) SetNillableContent(s *string) *CommentCreate {
	if s != nil {
		cc.SetContent(*s)
	}
	return cc
}

// SetPost sets the "post" edge to the Post entity.
func (cc *CommentCreate) SetPost(p *Post) *CommentCreate {
	return cc.SetPostID(p.ID)
}

// SetUser sets the "user" edge to the User entity.
func (cc *CommentCreate) SetUser(u *User) *CommentCreate {
	return cc.SetUserID(u.ID)
}

// Mutation returns the CommentMutation object of the builder.
func (cc *CommentCreate) Mutation() *CommentMutation {
	return cc.mutation
}

// Save creates the Comment in the database.
func (cc *CommentCreate) Save(ctx context.Context) (*Comment, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CommentCreate) SaveX(ctx context.Context) *Comment {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CommentCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CommentCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CommentCreate) defaults() {
	if _, ok := cc.mutation.CreateTime(); !ok {
		v := comment.DefaultCreateTime()
		cc.mutation.SetCreateTime(v)
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		v := comment.DefaultUpdateTime()
		cc.mutation.SetUpdateTime(v)
	}
	if _, ok := cc.mutation.UserID(); !ok {
		v := comment.DefaultUserID
		cc.mutation.SetUserID(v)
	}
	if _, ok := cc.mutation.PostID(); !ok {
		v := comment.DefaultPostID
		cc.mutation.SetPostID(v)
	}
	if _, ok := cc.mutation.Content(); !ok {
		v := comment.DefaultContent
		cc.mutation.SetContent(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CommentCreate) check() error {
	if _, ok := cc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Comment.create_time"`)}
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Comment.update_time"`)}
	}
	if _, ok := cc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Comment.content"`)}
	}
	return nil
}

func (cc *CommentCreate) sqlSave(ctx context.Context) (*Comment, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CommentCreate) createSpec() (*Comment, *sqlgraph.CreateSpec) {
	var (
		_node = &Comment{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(comment.Table, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt))
	)
	_spec.OnConflict = cc.conflict
	if value, ok := cc.mutation.CreateTime(); ok {
		_spec.SetField(comment.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := cc.mutation.UpdateTime(); ok {
		_spec.SetField(comment.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := cc.mutation.Content(); ok {
		_spec.SetField(comment.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if nodes := cc.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.PostTable,
			Columns: []string{comment.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PostID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.UserTable,
			Columns: []string{comment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Comment.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CommentUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (cc *CommentCreate) OnConflict(opts ...sql.ConflictOption) *CommentUpsertOne {
	cc.conflict = opts
	return &CommentUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Comment.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cc *CommentCreate) OnConflictColumns(columns ...string) *CommentUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &CommentUpsertOne{
		create: cc,
	}
}

type (
	// CommentUpsertOne is the builder for "upsert"-ing
	//  one Comment node.
	CommentUpsertOne struct {
		create *CommentCreate
	}

	// CommentUpsert is the "OnConflict" setter.
	CommentUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *CommentUpsert) SetUpdateTime(v time.Time) *CommentUpsert {
	u.Set(comment.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CommentUpsert) UpdateUpdateTime() *CommentUpsert {
	u.SetExcluded(comment.FieldUpdateTime)
	return u
}

// SetUserID sets the "user_id" field.
func (u *CommentUpsert) SetUserID(v int) *CommentUpsert {
	u.Set(comment.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *CommentUpsert) UpdateUserID() *CommentUpsert {
	u.SetExcluded(comment.FieldUserID)
	return u
}

// ClearUserID clears the value of the "user_id" field.
func (u *CommentUpsert) ClearUserID() *CommentUpsert {
	u.SetNull(comment.FieldUserID)
	return u
}

// SetPostID sets the "post_id" field.
func (u *CommentUpsert) SetPostID(v int) *CommentUpsert {
	u.Set(comment.FieldPostID, v)
	return u
}

// UpdatePostID sets the "post_id" field to the value that was provided on create.
func (u *CommentUpsert) UpdatePostID() *CommentUpsert {
	u.SetExcluded(comment.FieldPostID)
	return u
}

// ClearPostID clears the value of the "post_id" field.
func (u *CommentUpsert) ClearPostID() *CommentUpsert {
	u.SetNull(comment.FieldPostID)
	return u
}

// SetContent sets the "content" field.
func (u *CommentUpsert) SetContent(v string) *CommentUpsert {
	u.Set(comment.FieldContent, v)
	return u
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *CommentUpsert) UpdateContent() *CommentUpsert {
	u.SetExcluded(comment.FieldContent)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Comment.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *CommentUpsertOne) UpdateNewValues() *CommentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(comment.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Comment.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *CommentUpsertOne) Ignore() *CommentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CommentUpsertOne) DoNothing() *CommentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CommentCreate.OnConflict
// documentation for more info.
func (u *CommentUpsertOne) Update(set func(*CommentUpsert)) *CommentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CommentUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *CommentUpsertOne) SetUpdateTime(v time.Time) *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CommentUpsertOne) UpdateUpdateTime() *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetUserID sets the "user_id" field.
func (u *CommentUpsertOne) SetUserID(v int) *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *CommentUpsertOne) UpdateUserID() *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *CommentUpsertOne) ClearUserID() *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.ClearUserID()
	})
}

// SetPostID sets the "post_id" field.
func (u *CommentUpsertOne) SetPostID(v int) *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.SetPostID(v)
	})
}

// UpdatePostID sets the "post_id" field to the value that was provided on create.
func (u *CommentUpsertOne) UpdatePostID() *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.UpdatePostID()
	})
}

// ClearPostID clears the value of the "post_id" field.
func (u *CommentUpsertOne) ClearPostID() *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.ClearPostID()
	})
}

// SetContent sets the "content" field.
func (u *CommentUpsertOne) SetContent(v string) *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *CommentUpsertOne) UpdateContent() *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.UpdateContent()
	})
}

// Exec executes the query.
func (u *CommentUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CommentCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CommentUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CommentUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CommentUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CommentCreateBulk is the builder for creating many Comment entities in bulk.
type CommentCreateBulk struct {
	config
	err      error
	builders []*CommentCreate
	conflict []sql.ConflictOption
}

// Save creates the Comment entities in the database.
func (ccb *CommentCreateBulk) Save(ctx context.Context) ([]*Comment, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Comment, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommentMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CommentCreateBulk) SaveX(ctx context.Context) []*Comment {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CommentCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CommentCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Comment.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CommentUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (ccb *CommentCreateBulk) OnConflict(opts ...sql.ConflictOption) *CommentUpsertBulk {
	ccb.conflict = opts
	return &CommentUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Comment.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ccb *CommentCreateBulk) OnConflictColumns(columns ...string) *CommentUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &CommentUpsertBulk{
		create: ccb,
	}
}

// CommentUpsertBulk is the builder for "upsert"-ing
// a bulk of Comment nodes.
type CommentUpsertBulk struct {
	create *CommentCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Comment.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *CommentUpsertBulk) UpdateNewValues() *CommentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(comment.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Comment.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *CommentUpsertBulk) Ignore() *CommentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CommentUpsertBulk) DoNothing() *CommentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CommentCreateBulk.OnConflict
// documentation for more info.
func (u *CommentUpsertBulk) Update(set func(*CommentUpsert)) *CommentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CommentUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *CommentUpsertBulk) SetUpdateTime(v time.Time) *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CommentUpsertBulk) UpdateUpdateTime() *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetUserID sets the "user_id" field.
func (u *CommentUpsertBulk) SetUserID(v int) *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *CommentUpsertBulk) UpdateUserID() *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *CommentUpsertBulk) ClearUserID() *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.ClearUserID()
	})
}

// SetPostID sets the "post_id" field.
func (u *CommentUpsertBulk) SetPostID(v int) *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.SetPostID(v)
	})
}

// UpdatePostID sets the "post_id" field to the value that was provided on create.
func (u *CommentUpsertBulk) UpdatePostID() *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.UpdatePostID()
	})
}

// ClearPostID clears the value of the "post_id" field.
func (u *CommentUpsertBulk) ClearPostID() *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.ClearPostID()
	})
}

// SetContent sets the "content" field.
func (u *CommentUpsertBulk) SetContent(v string) *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *CommentUpsertBulk) UpdateContent() *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.UpdateContent()
	})
}

// Exec executes the query.
func (u *CommentUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CommentCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CommentCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CommentUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
