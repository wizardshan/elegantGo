// Code generated by ent, DO NOT EDIT.

package ent

import (
	"elegantGo/chapter-param-complex-validator-2/repository/ent/comment"
	"elegantGo/chapter-param-complex-validator-2/repository/ent/post"
	"elegantGo/chapter-param-complex-validator-2/repository/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PostCreate is the builder for creating a Post entity.
type PostCreate struct {
	config
	mutation *PostMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (pc *PostCreate) SetCreateTime(t time.Time) *PostCreate {
	pc.mutation.SetCreateTime(t)
	return pc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pc *PostCreate) SetNillableCreateTime(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetCreateTime(*t)
	}
	return pc
}

// SetUpdateTime sets the "update_time" field.
func (pc *PostCreate) SetUpdateTime(t time.Time) *PostCreate {
	pc.mutation.SetUpdateTime(t)
	return pc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pc *PostCreate) SetNillableUpdateTime(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetUpdateTime(*t)
	}
	return pc
}

// SetHashID sets the "hash_id" field.
func (pc *PostCreate) SetHashID(s string) *PostCreate {
	pc.mutation.SetHashID(s)
	return pc
}

// SetNillableHashID sets the "hash_id" field if the given value is not nil.
func (pc *PostCreate) SetNillableHashID(s *string) *PostCreate {
	if s != nil {
		pc.SetHashID(*s)
	}
	return pc
}

// SetUserID sets the "user_id" field.
func (pc *PostCreate) SetUserID(i int) *PostCreate {
	pc.mutation.SetUserID(i)
	return pc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (pc *PostCreate) SetNillableUserID(i *int) *PostCreate {
	if i != nil {
		pc.SetUserID(*i)
	}
	return pc
}

// SetTitle sets the "title" field.
func (pc *PostCreate) SetTitle(s string) *PostCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (pc *PostCreate) SetNillableTitle(s *string) *PostCreate {
	if s != nil {
		pc.SetTitle(*s)
	}
	return pc
}

// SetContent sets the "content" field.
func (pc *PostCreate) SetContent(s string) *PostCreate {
	pc.mutation.SetContent(s)
	return pc
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (pc *PostCreate) SetNillableContent(s *string) *PostCreate {
	if s != nil {
		pc.SetContent(*s)
	}
	return pc
}

// SetTimesOfRead sets the "times_of_read" field.
func (pc *PostCreate) SetTimesOfRead(i int) *PostCreate {
	pc.mutation.SetTimesOfRead(i)
	return pc
}

// SetNillableTimesOfRead sets the "times_of_read" field if the given value is not nil.
func (pc *PostCreate) SetNillableTimesOfRead(i *int) *PostCreate {
	if i != nil {
		pc.SetTimesOfRead(*i)
	}
	return pc
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (pc *PostCreate) AddCommentIDs(ids ...int) *PostCreate {
	pc.mutation.AddCommentIDs(ids...)
	return pc
}

// AddComments adds the "comments" edges to the Comment entity.
func (pc *PostCreate) AddComments(c ...*Comment) *PostCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddCommentIDs(ids...)
}

// SetUser sets the "user" edge to the User entity.
func (pc *PostCreate) SetUser(u *User) *PostCreate {
	return pc.SetUserID(u.ID)
}

// Mutation returns the PostMutation object of the builder.
func (pc *PostCreate) Mutation() *PostMutation {
	return pc.mutation
}

// Save creates the Post in the database.
func (pc *PostCreate) Save(ctx context.Context) (*Post, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PostCreate) SaveX(ctx context.Context) *Post {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PostCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PostCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PostCreate) defaults() {
	if _, ok := pc.mutation.CreateTime(); !ok {
		v := post.DefaultCreateTime()
		pc.mutation.SetCreateTime(v)
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		v := post.DefaultUpdateTime()
		pc.mutation.SetUpdateTime(v)
	}
	if _, ok := pc.mutation.HashID(); !ok {
		v := post.DefaultHashID
		pc.mutation.SetHashID(v)
	}
	if _, ok := pc.mutation.UserID(); !ok {
		v := post.DefaultUserID
		pc.mutation.SetUserID(v)
	}
	if _, ok := pc.mutation.Title(); !ok {
		v := post.DefaultTitle
		pc.mutation.SetTitle(v)
	}
	if _, ok := pc.mutation.Content(); !ok {
		v := post.DefaultContent
		pc.mutation.SetContent(v)
	}
	if _, ok := pc.mutation.TimesOfRead(); !ok {
		v := post.DefaultTimesOfRead
		pc.mutation.SetTimesOfRead(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PostCreate) check() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Post.create_time"`)}
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Post.update_time"`)}
	}
	if _, ok := pc.mutation.HashID(); !ok {
		return &ValidationError{Name: "hash_id", err: errors.New(`ent: missing required field "Post.hash_id"`)}
	}
	if _, ok := pc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Post.title"`)}
	}
	if _, ok := pc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Post.content"`)}
	}
	if _, ok := pc.mutation.TimesOfRead(); !ok {
		return &ValidationError{Name: "times_of_read", err: errors.New(`ent: missing required field "Post.times_of_read"`)}
	}
	return nil
}

func (pc *PostCreate) sqlSave(ctx context.Context) (*Post, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PostCreate) createSpec() (*Post, *sqlgraph.CreateSpec) {
	var (
		_node = &Post{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(post.Table, sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt))
	)
	_spec.OnConflict = pc.conflict
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.SetField(post.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.SetField(post.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := pc.mutation.HashID(); ok {
		_spec.SetField(post.FieldHashID, field.TypeString, value)
		_node.HashID = value
	}
	if value, ok := pc.mutation.Title(); ok {
		_spec.SetField(post.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := pc.mutation.Content(); ok {
		_spec.SetField(post.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := pc.mutation.TimesOfRead(); ok {
		_spec.SetField(post.FieldTimesOfRead, field.TypeInt, value)
		_node.TimesOfRead = value
	}
	if nodes := pc.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
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
	if nodes := pc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.UserTable,
			Columns: []string{post.UserColumn},
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
//	client.Post.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PostUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (pc *PostCreate) OnConflict(opts ...sql.ConflictOption) *PostUpsertOne {
	pc.conflict = opts
	return &PostUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Post.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *PostCreate) OnConflictColumns(columns ...string) *PostUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PostUpsertOne{
		create: pc,
	}
}

type (
	// PostUpsertOne is the builder for "upsert"-ing
	//  one Post node.
	PostUpsertOne struct {
		create *PostCreate
	}

	// PostUpsert is the "OnConflict" setter.
	PostUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *PostUpsert) SetUpdateTime(v time.Time) *PostUpsert {
	u.Set(post.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *PostUpsert) UpdateUpdateTime() *PostUpsert {
	u.SetExcluded(post.FieldUpdateTime)
	return u
}

// SetHashID sets the "hash_id" field.
func (u *PostUpsert) SetHashID(v string) *PostUpsert {
	u.Set(post.FieldHashID, v)
	return u
}

// UpdateHashID sets the "hash_id" field to the value that was provided on create.
func (u *PostUpsert) UpdateHashID() *PostUpsert {
	u.SetExcluded(post.FieldHashID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *PostUpsert) SetUserID(v int) *PostUpsert {
	u.Set(post.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *PostUpsert) UpdateUserID() *PostUpsert {
	u.SetExcluded(post.FieldUserID)
	return u
}

// ClearUserID clears the value of the "user_id" field.
func (u *PostUpsert) ClearUserID() *PostUpsert {
	u.SetNull(post.FieldUserID)
	return u
}

// SetTitle sets the "title" field.
func (u *PostUpsert) SetTitle(v string) *PostUpsert {
	u.Set(post.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *PostUpsert) UpdateTitle() *PostUpsert {
	u.SetExcluded(post.FieldTitle)
	return u
}

// SetContent sets the "content" field.
func (u *PostUpsert) SetContent(v string) *PostUpsert {
	u.Set(post.FieldContent, v)
	return u
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *PostUpsert) UpdateContent() *PostUpsert {
	u.SetExcluded(post.FieldContent)
	return u
}

// SetTimesOfRead sets the "times_of_read" field.
func (u *PostUpsert) SetTimesOfRead(v int) *PostUpsert {
	u.Set(post.FieldTimesOfRead, v)
	return u
}

// UpdateTimesOfRead sets the "times_of_read" field to the value that was provided on create.
func (u *PostUpsert) UpdateTimesOfRead() *PostUpsert {
	u.SetExcluded(post.FieldTimesOfRead)
	return u
}

// AddTimesOfRead adds v to the "times_of_read" field.
func (u *PostUpsert) AddTimesOfRead(v int) *PostUpsert {
	u.Add(post.FieldTimesOfRead, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Post.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *PostUpsertOne) UpdateNewValues() *PostUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(post.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Post.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PostUpsertOne) Ignore() *PostUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PostUpsertOne) DoNothing() *PostUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PostCreate.OnConflict
// documentation for more info.
func (u *PostUpsertOne) Update(set func(*PostUpsert)) *PostUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PostUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *PostUpsertOne) SetUpdateTime(v time.Time) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *PostUpsertOne) UpdateUpdateTime() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetHashID sets the "hash_id" field.
func (u *PostUpsertOne) SetHashID(v string) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetHashID(v)
	})
}

// UpdateHashID sets the "hash_id" field to the value that was provided on create.
func (u *PostUpsertOne) UpdateHashID() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdateHashID()
	})
}

// SetUserID sets the "user_id" field.
func (u *PostUpsertOne) SetUserID(v int) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *PostUpsertOne) UpdateUserID() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *PostUpsertOne) ClearUserID() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.ClearUserID()
	})
}

// SetTitle sets the "title" field.
func (u *PostUpsertOne) SetTitle(v string) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *PostUpsertOne) UpdateTitle() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdateTitle()
	})
}

// SetContent sets the "content" field.
func (u *PostUpsertOne) SetContent(v string) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *PostUpsertOne) UpdateContent() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdateContent()
	})
}

// SetTimesOfRead sets the "times_of_read" field.
func (u *PostUpsertOne) SetTimesOfRead(v int) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetTimesOfRead(v)
	})
}

// AddTimesOfRead adds v to the "times_of_read" field.
func (u *PostUpsertOne) AddTimesOfRead(v int) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.AddTimesOfRead(v)
	})
}

// UpdateTimesOfRead sets the "times_of_read" field to the value that was provided on create.
func (u *PostUpsertOne) UpdateTimesOfRead() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdateTimesOfRead()
	})
}

// Exec executes the query.
func (u *PostUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PostCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PostUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PostUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PostUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PostCreateBulk is the builder for creating many Post entities in bulk.
type PostCreateBulk struct {
	config
	err      error
	builders []*PostCreate
	conflict []sql.ConflictOption
}

// Save creates the Post entities in the database.
func (pcb *PostCreateBulk) Save(ctx context.Context) ([]*Post, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Post, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PostMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PostCreateBulk) SaveX(ctx context.Context) []*Post {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PostCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PostCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Post.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PostUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (pcb *PostCreateBulk) OnConflict(opts ...sql.ConflictOption) *PostUpsertBulk {
	pcb.conflict = opts
	return &PostUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Post.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *PostCreateBulk) OnConflictColumns(columns ...string) *PostUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PostUpsertBulk{
		create: pcb,
	}
}

// PostUpsertBulk is the builder for "upsert"-ing
// a bulk of Post nodes.
type PostUpsertBulk struct {
	create *PostCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Post.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *PostUpsertBulk) UpdateNewValues() *PostUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(post.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Post.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PostUpsertBulk) Ignore() *PostUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PostUpsertBulk) DoNothing() *PostUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PostCreateBulk.OnConflict
// documentation for more info.
func (u *PostUpsertBulk) Update(set func(*PostUpsert)) *PostUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PostUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *PostUpsertBulk) SetUpdateTime(v time.Time) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdateUpdateTime() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetHashID sets the "hash_id" field.
func (u *PostUpsertBulk) SetHashID(v string) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetHashID(v)
	})
}

// UpdateHashID sets the "hash_id" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdateHashID() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdateHashID()
	})
}

// SetUserID sets the "user_id" field.
func (u *PostUpsertBulk) SetUserID(v int) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdateUserID() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *PostUpsertBulk) ClearUserID() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.ClearUserID()
	})
}

// SetTitle sets the "title" field.
func (u *PostUpsertBulk) SetTitle(v string) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdateTitle() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdateTitle()
	})
}

// SetContent sets the "content" field.
func (u *PostUpsertBulk) SetContent(v string) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdateContent() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdateContent()
	})
}

// SetTimesOfRead sets the "times_of_read" field.
func (u *PostUpsertBulk) SetTimesOfRead(v int) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetTimesOfRead(v)
	})
}

// AddTimesOfRead adds v to the "times_of_read" field.
func (u *PostUpsertBulk) AddTimesOfRead(v int) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.AddTimesOfRead(v)
	})
}

// UpdateTimesOfRead sets the "times_of_read" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdateTimesOfRead() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdateTimesOfRead()
	})
}

// Exec executes the query.
func (u *PostUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PostCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PostCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PostUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}