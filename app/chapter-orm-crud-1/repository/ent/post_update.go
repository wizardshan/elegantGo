// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/chapter-orm-crud-1/repository/ent/comment"
	"app/chapter-orm-crud-1/repository/ent/post"
	"app/chapter-orm-crud-1/repository/ent/predicate"
	"app/chapter-orm-crud-1/repository/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PostUpdate is the builder for updating Post entities.
type PostUpdate struct {
	config
	hooks    []Hook
	mutation *PostMutation
}

// Where appends a list predicates to the PostUpdate builder.
func (pu *PostUpdate) Where(ps ...predicate.Post) *PostUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdateTime sets the "update_time" field.
func (pu *PostUpdate) SetUpdateTime(t time.Time) *PostUpdate {
	pu.mutation.SetUpdateTime(t)
	return pu
}

// SetHashID sets the "hash_id" field.
func (pu *PostUpdate) SetHashID(s string) *PostUpdate {
	pu.mutation.SetHashID(s)
	return pu
}

// SetNillableHashID sets the "hash_id" field if the given value is not nil.
func (pu *PostUpdate) SetNillableHashID(s *string) *PostUpdate {
	if s != nil {
		pu.SetHashID(*s)
	}
	return pu
}

// SetUserID sets the "user_id" field.
func (pu *PostUpdate) SetUserID(i int) *PostUpdate {
	pu.mutation.SetUserID(i)
	return pu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (pu *PostUpdate) SetNillableUserID(i *int) *PostUpdate {
	if i != nil {
		pu.SetUserID(*i)
	}
	return pu
}

// ClearUserID clears the value of the "user_id" field.
func (pu *PostUpdate) ClearUserID() *PostUpdate {
	pu.mutation.ClearUserID()
	return pu
}

// SetTitle sets the "title" field.
func (pu *PostUpdate) SetTitle(s string) *PostUpdate {
	pu.mutation.SetTitle(s)
	return pu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (pu *PostUpdate) SetNillableTitle(s *string) *PostUpdate {
	if s != nil {
		pu.SetTitle(*s)
	}
	return pu
}

// SetContent sets the "content" field.
func (pu *PostUpdate) SetContent(s string) *PostUpdate {
	pu.mutation.SetContent(s)
	return pu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (pu *PostUpdate) SetNillableContent(s *string) *PostUpdate {
	if s != nil {
		pu.SetContent(*s)
	}
	return pu
}

// SetTimesOfRead sets the "times_of_read" field.
func (pu *PostUpdate) SetTimesOfRead(i int) *PostUpdate {
	pu.mutation.ResetTimesOfRead()
	pu.mutation.SetTimesOfRead(i)
	return pu
}

// SetNillableTimesOfRead sets the "times_of_read" field if the given value is not nil.
func (pu *PostUpdate) SetNillableTimesOfRead(i *int) *PostUpdate {
	if i != nil {
		pu.SetTimesOfRead(*i)
	}
	return pu
}

// AddTimesOfRead adds i to the "times_of_read" field.
func (pu *PostUpdate) AddTimesOfRead(i int) *PostUpdate {
	pu.mutation.AddTimesOfRead(i)
	return pu
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (pu *PostUpdate) AddCommentIDs(ids ...int) *PostUpdate {
	pu.mutation.AddCommentIDs(ids...)
	return pu
}

// AddComments adds the "comments" edges to the Comment entity.
func (pu *PostUpdate) AddComments(c ...*Comment) *PostUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.AddCommentIDs(ids...)
}

// SetUser sets the "user" edge to the User entity.
func (pu *PostUpdate) SetUser(u *User) *PostUpdate {
	return pu.SetUserID(u.ID)
}

// Mutation returns the PostMutation object of the builder.
func (pu *PostUpdate) Mutation() *PostMutation {
	return pu.mutation
}

// ClearComments clears all "comments" edges to the Comment entity.
func (pu *PostUpdate) ClearComments() *PostUpdate {
	pu.mutation.ClearComments()
	return pu
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (pu *PostUpdate) RemoveCommentIDs(ids ...int) *PostUpdate {
	pu.mutation.RemoveCommentIDs(ids...)
	return pu
}

// RemoveComments removes "comments" edges to Comment entities.
func (pu *PostUpdate) RemoveComments(c ...*Comment) *PostUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.RemoveCommentIDs(ids...)
}

// ClearUser clears the "user" edge to the User entity.
func (pu *PostUpdate) ClearUser() *PostUpdate {
	pu.mutation.ClearUser()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PostUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PostUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PostUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PostUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PostUpdate) defaults() {
	if _, ok := pu.mutation.UpdateTime(); !ok {
		v := post.UpdateDefaultUpdateTime()
		pu.mutation.SetUpdateTime(v)
	}
}

func (pu *PostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(post.Table, post.Columns, sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdateTime(); ok {
		_spec.SetField(post.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := pu.mutation.HashID(); ok {
		_spec.SetField(post.FieldHashID, field.TypeString, value)
	}
	if value, ok := pu.mutation.Title(); ok {
		_spec.SetField(post.FieldTitle, field.TypeString, value)
	}
	if value, ok := pu.mutation.Content(); ok {
		_spec.SetField(post.FieldContent, field.TypeString, value)
	}
	if value, ok := pu.mutation.TimesOfRead(); ok {
		_spec.SetField(post.FieldTimesOfRead, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedTimesOfRead(); ok {
		_spec.AddField(post.FieldTimesOfRead, field.TypeInt, value)
	}
	if pu.mutation.CommentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !pu.mutation.CommentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.CommentsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PostUpdateOne is the builder for updating a single Post entity.
type PostUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PostMutation
}

// SetUpdateTime sets the "update_time" field.
func (puo *PostUpdateOne) SetUpdateTime(t time.Time) *PostUpdateOne {
	puo.mutation.SetUpdateTime(t)
	return puo
}

// SetHashID sets the "hash_id" field.
func (puo *PostUpdateOne) SetHashID(s string) *PostUpdateOne {
	puo.mutation.SetHashID(s)
	return puo
}

// SetNillableHashID sets the "hash_id" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableHashID(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetHashID(*s)
	}
	return puo
}

// SetUserID sets the "user_id" field.
func (puo *PostUpdateOne) SetUserID(i int) *PostUpdateOne {
	puo.mutation.SetUserID(i)
	return puo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableUserID(i *int) *PostUpdateOne {
	if i != nil {
		puo.SetUserID(*i)
	}
	return puo
}

// ClearUserID clears the value of the "user_id" field.
func (puo *PostUpdateOne) ClearUserID() *PostUpdateOne {
	puo.mutation.ClearUserID()
	return puo
}

// SetTitle sets the "title" field.
func (puo *PostUpdateOne) SetTitle(s string) *PostUpdateOne {
	puo.mutation.SetTitle(s)
	return puo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableTitle(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetTitle(*s)
	}
	return puo
}

// SetContent sets the "content" field.
func (puo *PostUpdateOne) SetContent(s string) *PostUpdateOne {
	puo.mutation.SetContent(s)
	return puo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableContent(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetContent(*s)
	}
	return puo
}

// SetTimesOfRead sets the "times_of_read" field.
func (puo *PostUpdateOne) SetTimesOfRead(i int) *PostUpdateOne {
	puo.mutation.ResetTimesOfRead()
	puo.mutation.SetTimesOfRead(i)
	return puo
}

// SetNillableTimesOfRead sets the "times_of_read" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableTimesOfRead(i *int) *PostUpdateOne {
	if i != nil {
		puo.SetTimesOfRead(*i)
	}
	return puo
}

// AddTimesOfRead adds i to the "times_of_read" field.
func (puo *PostUpdateOne) AddTimesOfRead(i int) *PostUpdateOne {
	puo.mutation.AddTimesOfRead(i)
	return puo
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (puo *PostUpdateOne) AddCommentIDs(ids ...int) *PostUpdateOne {
	puo.mutation.AddCommentIDs(ids...)
	return puo
}

// AddComments adds the "comments" edges to the Comment entity.
func (puo *PostUpdateOne) AddComments(c ...*Comment) *PostUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.AddCommentIDs(ids...)
}

// SetUser sets the "user" edge to the User entity.
func (puo *PostUpdateOne) SetUser(u *User) *PostUpdateOne {
	return puo.SetUserID(u.ID)
}

// Mutation returns the PostMutation object of the builder.
func (puo *PostUpdateOne) Mutation() *PostMutation {
	return puo.mutation
}

// ClearComments clears all "comments" edges to the Comment entity.
func (puo *PostUpdateOne) ClearComments() *PostUpdateOne {
	puo.mutation.ClearComments()
	return puo
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (puo *PostUpdateOne) RemoveCommentIDs(ids ...int) *PostUpdateOne {
	puo.mutation.RemoveCommentIDs(ids...)
	return puo
}

// RemoveComments removes "comments" edges to Comment entities.
func (puo *PostUpdateOne) RemoveComments(c ...*Comment) *PostUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.RemoveCommentIDs(ids...)
}

// ClearUser clears the "user" edge to the User entity.
func (puo *PostUpdateOne) ClearUser() *PostUpdateOne {
	puo.mutation.ClearUser()
	return puo
}

// Where appends a list predicates to the PostUpdate builder.
func (puo *PostUpdateOne) Where(ps ...predicate.Post) *PostUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PostUpdateOne) Select(field string, fields ...string) *PostUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Post entity.
func (puo *PostUpdateOne) Save(ctx context.Context) (*Post, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PostUpdateOne) SaveX(ctx context.Context) *Post {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PostUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PostUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PostUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdateTime(); !ok {
		v := post.UpdateDefaultUpdateTime()
		puo.mutation.SetUpdateTime(v)
	}
}

func (puo *PostUpdateOne) sqlSave(ctx context.Context) (_node *Post, err error) {
	_spec := sqlgraph.NewUpdateSpec(post.Table, post.Columns, sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Post.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, post.FieldID)
		for _, f := range fields {
			if !post.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != post.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdateTime(); ok {
		_spec.SetField(post.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := puo.mutation.HashID(); ok {
		_spec.SetField(post.FieldHashID, field.TypeString, value)
	}
	if value, ok := puo.mutation.Title(); ok {
		_spec.SetField(post.FieldTitle, field.TypeString, value)
	}
	if value, ok := puo.mutation.Content(); ok {
		_spec.SetField(post.FieldContent, field.TypeString, value)
	}
	if value, ok := puo.mutation.TimesOfRead(); ok {
		_spec.SetField(post.FieldTimesOfRead, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedTimesOfRead(); ok {
		_spec.AddField(post.FieldTimesOfRead, field.TypeInt, value)
	}
	if puo.mutation.CommentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !puo.mutation.CommentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.CommentsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Post{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
