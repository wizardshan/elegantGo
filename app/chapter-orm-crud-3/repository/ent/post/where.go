// Code generated by ent, DO NOT EDIT.

package post

import (
	"app/chapter-orm-crud-3/repository/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldUpdateTime, v))
}

// HashID applies equality check predicate on the "hash_id" field. It's identical to HashIDEQ.
func HashID(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldHashID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldUserID, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldTitle, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldContent, v))
}

// TimesOfRead applies equality check predicate on the "times_of_read" field. It's identical to TimesOfReadEQ.
func TimesOfRead(v int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldTimesOfRead, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldUpdateTime, v))
}

// HashIDEQ applies the EQ predicate on the "hash_id" field.
func HashIDEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldHashID, v))
}

// HashIDNEQ applies the NEQ predicate on the "hash_id" field.
func HashIDNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldHashID, v))
}

// HashIDIn applies the In predicate on the "hash_id" field.
func HashIDIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldHashID, vs...))
}

// HashIDNotIn applies the NotIn predicate on the "hash_id" field.
func HashIDNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldHashID, vs...))
}

// HashIDGT applies the GT predicate on the "hash_id" field.
func HashIDGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldHashID, v))
}

// HashIDGTE applies the GTE predicate on the "hash_id" field.
func HashIDGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldHashID, v))
}

// HashIDLT applies the LT predicate on the "hash_id" field.
func HashIDLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldHashID, v))
}

// HashIDLTE applies the LTE predicate on the "hash_id" field.
func HashIDLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldHashID, v))
}

// HashIDContains applies the Contains predicate on the "hash_id" field.
func HashIDContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldHashID, v))
}

// HashIDHasPrefix applies the HasPrefix predicate on the "hash_id" field.
func HashIDHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldHashID, v))
}

// HashIDHasSuffix applies the HasSuffix predicate on the "hash_id" field.
func HashIDHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldHashID, v))
}

// HashIDEqualFold applies the EqualFold predicate on the "hash_id" field.
func HashIDEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldHashID, v))
}

// HashIDContainsFold applies the ContainsFold predicate on the "hash_id" field.
func HashIDContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldHashID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Post {
	return predicate.Post(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Post {
	return predicate.Post(sql.FieldNotNull(FieldUserID))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldTitle, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldContent, v))
}

// TimesOfReadEQ applies the EQ predicate on the "times_of_read" field.
func TimesOfReadEQ(v int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldTimesOfRead, v))
}

// TimesOfReadNEQ applies the NEQ predicate on the "times_of_read" field.
func TimesOfReadNEQ(v int) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldTimesOfRead, v))
}

// TimesOfReadIn applies the In predicate on the "times_of_read" field.
func TimesOfReadIn(vs ...int) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldTimesOfRead, vs...))
}

// TimesOfReadNotIn applies the NotIn predicate on the "times_of_read" field.
func TimesOfReadNotIn(vs ...int) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldTimesOfRead, vs...))
}

// TimesOfReadGT applies the GT predicate on the "times_of_read" field.
func TimesOfReadGT(v int) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldTimesOfRead, v))
}

// TimesOfReadGTE applies the GTE predicate on the "times_of_read" field.
func TimesOfReadGTE(v int) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldTimesOfRead, v))
}

// TimesOfReadLT applies the LT predicate on the "times_of_read" field.
func TimesOfReadLT(v int) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldTimesOfRead, v))
}

// TimesOfReadLTE applies the LTE predicate on the "times_of_read" field.
func TimesOfReadLTE(v int) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldTimesOfRead, v))
}

// HasComments applies the HasEdge predicate on the "comments" edge.
func HasComments() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CommentsTable, CommentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentsWith applies the HasEdge predicate on the "comments" edge with a given conditions (other predicates).
func HasCommentsWith(preds ...predicate.Comment) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := newCommentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Post) predicate.Post {
	return predicate.Post(sql.NotPredicates(p))
}
