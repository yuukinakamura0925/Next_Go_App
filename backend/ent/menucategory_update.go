// Code generated by ent, DO NOT EDIT.

package ent

import (
	"Next_Go_App/ent/menucategory"
	"Next_Go_App/ent/predicate"
	"Next_Go_App/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MenuCategoryUpdate is the builder for updating MenuCategory entities.
type MenuCategoryUpdate struct {
	config
	hooks    []Hook
	mutation *MenuCategoryMutation
}

// Where appends a list predicates to the MenuCategoryUpdate builder.
func (mcu *MenuCategoryUpdate) Where(ps ...predicate.MenuCategory) *MenuCategoryUpdate {
	mcu.mutation.Where(ps...)
	return mcu
}

// SetUserID sets the "user_id" field.
func (mcu *MenuCategoryUpdate) SetUserID(i int) *MenuCategoryUpdate {
	mcu.mutation.SetUserID(i)
	return mcu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (mcu *MenuCategoryUpdate) SetNillableUserID(i *int) *MenuCategoryUpdate {
	if i != nil {
		mcu.SetUserID(*i)
	}
	return mcu
}

// SetName sets the "name" field.
func (mcu *MenuCategoryUpdate) SetName(s string) *MenuCategoryUpdate {
	mcu.mutation.SetName(s)
	return mcu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mcu *MenuCategoryUpdate) SetNillableName(s *string) *MenuCategoryUpdate {
	if s != nil {
		mcu.SetName(*s)
	}
	return mcu
}

// SetCreatedAt sets the "created_at" field.
func (mcu *MenuCategoryUpdate) SetCreatedAt(t time.Time) *MenuCategoryUpdate {
	mcu.mutation.SetCreatedAt(t)
	return mcu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mcu *MenuCategoryUpdate) SetNillableCreatedAt(t *time.Time) *MenuCategoryUpdate {
	if t != nil {
		mcu.SetCreatedAt(*t)
	}
	return mcu
}

// SetUpdatedAt sets the "updated_at" field.
func (mcu *MenuCategoryUpdate) SetUpdatedAt(t time.Time) *MenuCategoryUpdate {
	mcu.mutation.SetUpdatedAt(t)
	return mcu
}

// SetUser sets the "user" edge to the User entity.
func (mcu *MenuCategoryUpdate) SetUser(u *User) *MenuCategoryUpdate {
	return mcu.SetUserID(u.ID)
}

// Mutation returns the MenuCategoryMutation object of the builder.
func (mcu *MenuCategoryUpdate) Mutation() *MenuCategoryMutation {
	return mcu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (mcu *MenuCategoryUpdate) ClearUser() *MenuCategoryUpdate {
	mcu.mutation.ClearUser()
	return mcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mcu *MenuCategoryUpdate) Save(ctx context.Context) (int, error) {
	mcu.defaults()
	return withHooks(ctx, mcu.sqlSave, mcu.mutation, mcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mcu *MenuCategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := mcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mcu *MenuCategoryUpdate) Exec(ctx context.Context) error {
	_, err := mcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcu *MenuCategoryUpdate) ExecX(ctx context.Context) {
	if err := mcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mcu *MenuCategoryUpdate) defaults() {
	if _, ok := mcu.mutation.UpdatedAt(); !ok {
		v := menucategory.UpdateDefaultUpdatedAt()
		mcu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mcu *MenuCategoryUpdate) check() error {
	if v, ok := mcu.mutation.UserID(); ok {
		if err := menucategory.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "MenuCategory.user_id": %w`, err)}
		}
	}
	if v, ok := mcu.mutation.Name(); ok {
		if err := menucategory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "MenuCategory.name": %w`, err)}
		}
	}
	if mcu.mutation.UserCleared() && len(mcu.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "MenuCategory.user"`)
	}
	return nil
}

func (mcu *MenuCategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(menucategory.Table, menucategory.Columns, sqlgraph.NewFieldSpec(menucategory.FieldID, field.TypeInt))
	if ps := mcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mcu.mutation.Name(); ok {
		_spec.SetField(menucategory.FieldName, field.TypeString, value)
	}
	if value, ok := mcu.mutation.CreatedAt(); ok {
		_spec.SetField(menucategory.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := mcu.mutation.UpdatedAt(); ok {
		_spec.SetField(menucategory.FieldUpdatedAt, field.TypeTime, value)
	}
	if mcu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menucategory.UserTable,
			Columns: []string{menucategory.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menucategory.UserTable,
			Columns: []string{menucategory.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, mcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{menucategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mcu.mutation.done = true
	return n, nil
}

// MenuCategoryUpdateOne is the builder for updating a single MenuCategory entity.
type MenuCategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MenuCategoryMutation
}

// SetUserID sets the "user_id" field.
func (mcuo *MenuCategoryUpdateOne) SetUserID(i int) *MenuCategoryUpdateOne {
	mcuo.mutation.SetUserID(i)
	return mcuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (mcuo *MenuCategoryUpdateOne) SetNillableUserID(i *int) *MenuCategoryUpdateOne {
	if i != nil {
		mcuo.SetUserID(*i)
	}
	return mcuo
}

// SetName sets the "name" field.
func (mcuo *MenuCategoryUpdateOne) SetName(s string) *MenuCategoryUpdateOne {
	mcuo.mutation.SetName(s)
	return mcuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mcuo *MenuCategoryUpdateOne) SetNillableName(s *string) *MenuCategoryUpdateOne {
	if s != nil {
		mcuo.SetName(*s)
	}
	return mcuo
}

// SetCreatedAt sets the "created_at" field.
func (mcuo *MenuCategoryUpdateOne) SetCreatedAt(t time.Time) *MenuCategoryUpdateOne {
	mcuo.mutation.SetCreatedAt(t)
	return mcuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mcuo *MenuCategoryUpdateOne) SetNillableCreatedAt(t *time.Time) *MenuCategoryUpdateOne {
	if t != nil {
		mcuo.SetCreatedAt(*t)
	}
	return mcuo
}

// SetUpdatedAt sets the "updated_at" field.
func (mcuo *MenuCategoryUpdateOne) SetUpdatedAt(t time.Time) *MenuCategoryUpdateOne {
	mcuo.mutation.SetUpdatedAt(t)
	return mcuo
}

// SetUser sets the "user" edge to the User entity.
func (mcuo *MenuCategoryUpdateOne) SetUser(u *User) *MenuCategoryUpdateOne {
	return mcuo.SetUserID(u.ID)
}

// Mutation returns the MenuCategoryMutation object of the builder.
func (mcuo *MenuCategoryUpdateOne) Mutation() *MenuCategoryMutation {
	return mcuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (mcuo *MenuCategoryUpdateOne) ClearUser() *MenuCategoryUpdateOne {
	mcuo.mutation.ClearUser()
	return mcuo
}

// Where appends a list predicates to the MenuCategoryUpdate builder.
func (mcuo *MenuCategoryUpdateOne) Where(ps ...predicate.MenuCategory) *MenuCategoryUpdateOne {
	mcuo.mutation.Where(ps...)
	return mcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mcuo *MenuCategoryUpdateOne) Select(field string, fields ...string) *MenuCategoryUpdateOne {
	mcuo.fields = append([]string{field}, fields...)
	return mcuo
}

// Save executes the query and returns the updated MenuCategory entity.
func (mcuo *MenuCategoryUpdateOne) Save(ctx context.Context) (*MenuCategory, error) {
	mcuo.defaults()
	return withHooks(ctx, mcuo.sqlSave, mcuo.mutation, mcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mcuo *MenuCategoryUpdateOne) SaveX(ctx context.Context) *MenuCategory {
	node, err := mcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mcuo *MenuCategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := mcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcuo *MenuCategoryUpdateOne) ExecX(ctx context.Context) {
	if err := mcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mcuo *MenuCategoryUpdateOne) defaults() {
	if _, ok := mcuo.mutation.UpdatedAt(); !ok {
		v := menucategory.UpdateDefaultUpdatedAt()
		mcuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mcuo *MenuCategoryUpdateOne) check() error {
	if v, ok := mcuo.mutation.UserID(); ok {
		if err := menucategory.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "MenuCategory.user_id": %w`, err)}
		}
	}
	if v, ok := mcuo.mutation.Name(); ok {
		if err := menucategory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "MenuCategory.name": %w`, err)}
		}
	}
	if mcuo.mutation.UserCleared() && len(mcuo.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "MenuCategory.user"`)
	}
	return nil
}

func (mcuo *MenuCategoryUpdateOne) sqlSave(ctx context.Context) (_node *MenuCategory, err error) {
	if err := mcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(menucategory.Table, menucategory.Columns, sqlgraph.NewFieldSpec(menucategory.FieldID, field.TypeInt))
	id, ok := mcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MenuCategory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, menucategory.FieldID)
		for _, f := range fields {
			if !menucategory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != menucategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mcuo.mutation.Name(); ok {
		_spec.SetField(menucategory.FieldName, field.TypeString, value)
	}
	if value, ok := mcuo.mutation.CreatedAt(); ok {
		_spec.SetField(menucategory.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := mcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(menucategory.FieldUpdatedAt, field.TypeTime, value)
	}
	if mcuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menucategory.UserTable,
			Columns: []string{menucategory.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menucategory.UserTable,
			Columns: []string{menucategory.UserColumn},
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
	_node = &MenuCategory{config: mcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{menucategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mcuo.mutation.done = true
	return _node, nil
}
