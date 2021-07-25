// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/jackc/pgtype"
	"github.com/stackworx-go/entext/internal/integration/ent/predicate"
	"github.com/stackworx-go/entext/internal/integration/ent/user"
	"github.com/stackworx-go/entext/internal/integration/ent/userstatus"
)

// UserStatusUpdate is the builder for updating UserStatus entities.
type UserStatusUpdate struct {
	config
	hooks    []Hook
	mutation *UserStatusMutation
}

// Where appends a list predicates to the UserStatusUpdate builder.
func (usu *UserStatusUpdate) Where(ps ...predicate.UserStatus) *UserStatusUpdate {
	usu.mutation.Where(ps...)
	return usu
}

// SetCreatedByID sets the "created_by_id" field.
func (usu *UserStatusUpdate) SetCreatedByID(i int) *UserStatusUpdate {
	usu.mutation.SetCreatedByID(i)
	return usu
}

// SetUpdatedAt sets the "updated_at" field.
func (usu *UserStatusUpdate) SetUpdatedAt(t time.Time) *UserStatusUpdate {
	usu.mutation.SetUpdatedAt(t)
	return usu
}

// SetUpdatedByID sets the "updated_by_id" field.
func (usu *UserStatusUpdate) SetUpdatedByID(i int) *UserStatusUpdate {
	usu.mutation.SetUpdatedByID(i)
	return usu
}

// SetNillableUpdatedByID sets the "updated_by_id" field if the given value is not nil.
func (usu *UserStatusUpdate) SetNillableUpdatedByID(i *int) *UserStatusUpdate {
	if i != nil {
		usu.SetUpdatedByID(*i)
	}
	return usu
}

// ClearUpdatedByID clears the value of the "updated_by_id" field.
func (usu *UserStatusUpdate) ClearUpdatedByID() *UserStatusUpdate {
	usu.mutation.ClearUpdatedByID()
	return usu
}

// SetActive sets the "active" field.
func (usu *UserStatusUpdate) SetActive(b bool) *UserStatusUpdate {
	usu.mutation.SetActive(b)
	return usu
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (usu *UserStatusUpdate) SetNillableActive(b *bool) *UserStatusUpdate {
	if b != nil {
		usu.SetActive(*b)
	}
	return usu
}

// SetDuration sets the "duration" field.
func (usu *UserStatusUpdate) SetDuration(pg *pgtype.Tstzrange) *UserStatusUpdate {
	usu.mutation.SetDuration(pg)
	return usu
}

// ClearDuration clears the value of the "duration" field.
func (usu *UserStatusUpdate) ClearDuration() *UserStatusUpdate {
	usu.mutation.ClearDuration()
	return usu
}

// SetUserID sets the "user_id" field.
func (usu *UserStatusUpdate) SetUserID(i int) *UserStatusUpdate {
	usu.mutation.SetUserID(i)
	return usu
}

// SetCreatedBy sets the "created_by" edge to the User entity.
func (usu *UserStatusUpdate) SetCreatedBy(u *User) *UserStatusUpdate {
	return usu.SetCreatedByID(u.ID)
}

// SetUpdatedBy sets the "updated_by" edge to the User entity.
func (usu *UserStatusUpdate) SetUpdatedBy(u *User) *UserStatusUpdate {
	return usu.SetUpdatedByID(u.ID)
}

// SetUser sets the "user" edge to the User entity.
func (usu *UserStatusUpdate) SetUser(u *User) *UserStatusUpdate {
	return usu.SetUserID(u.ID)
}

// Mutation returns the UserStatusMutation object of the builder.
func (usu *UserStatusUpdate) Mutation() *UserStatusMutation {
	return usu.mutation
}

// ClearCreatedBy clears the "created_by" edge to the User entity.
func (usu *UserStatusUpdate) ClearCreatedBy() *UserStatusUpdate {
	usu.mutation.ClearCreatedBy()
	return usu
}

// ClearUpdatedBy clears the "updated_by" edge to the User entity.
func (usu *UserStatusUpdate) ClearUpdatedBy() *UserStatusUpdate {
	usu.mutation.ClearUpdatedBy()
	return usu
}

// ClearUser clears the "user" edge to the User entity.
func (usu *UserStatusUpdate) ClearUser() *UserStatusUpdate {
	usu.mutation.ClearUser()
	return usu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (usu *UserStatusUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := usu.defaults(); err != nil {
		return 0, err
	}
	if len(usu.hooks) == 0 {
		if err = usu.check(); err != nil {
			return 0, err
		}
		affected, err = usu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = usu.check(); err != nil {
				return 0, err
			}
			usu.mutation = mutation
			affected, err = usu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(usu.hooks) - 1; i >= 0; i-- {
			if usu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = usu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, usu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (usu *UserStatusUpdate) SaveX(ctx context.Context) int {
	affected, err := usu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (usu *UserStatusUpdate) Exec(ctx context.Context) error {
	_, err := usu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usu *UserStatusUpdate) ExecX(ctx context.Context) {
	if err := usu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (usu *UserStatusUpdate) defaults() error {
	if _, ok := usu.mutation.UpdatedAt(); !ok {
		if userstatus.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized userstatus.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := userstatus.UpdateDefaultUpdatedAt()
		usu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (usu *UserStatusUpdate) check() error {
	if _, ok := usu.mutation.CreatedByID(); usu.mutation.CreatedByCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"created_by\"")
	}
	if _, ok := usu.mutation.UserID(); usu.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (usu *UserStatusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userstatus.Table,
			Columns: userstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userstatus.FieldID,
			},
		},
	}
	if ps := usu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := usu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userstatus.FieldUpdatedAt,
		})
	}
	if value, ok := usu.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: userstatus.FieldActive,
		})
	}
	if value, ok := usu.mutation.Duration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: userstatus.FieldDuration,
		})
	}
	if usu.mutation.DurationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: userstatus.FieldDuration,
		})
	}
	if usu.mutation.CreatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userstatus.CreatedByTable,
			Columns: []string{userstatus.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := usu.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userstatus.CreatedByTable,
			Columns: []string{userstatus.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if usu.mutation.UpdatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userstatus.UpdatedByTable,
			Columns: []string{userstatus.UpdatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := usu.mutation.UpdatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userstatus.UpdatedByTable,
			Columns: []string{userstatus.UpdatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if usu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userstatus.UserTable,
			Columns: []string{userstatus.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := usu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userstatus.UserTable,
			Columns: []string{userstatus.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, usu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userstatus.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserStatusUpdateOne is the builder for updating a single UserStatus entity.
type UserStatusUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserStatusMutation
}

// SetCreatedByID sets the "created_by_id" field.
func (usuo *UserStatusUpdateOne) SetCreatedByID(i int) *UserStatusUpdateOne {
	usuo.mutation.SetCreatedByID(i)
	return usuo
}

// SetUpdatedAt sets the "updated_at" field.
func (usuo *UserStatusUpdateOne) SetUpdatedAt(t time.Time) *UserStatusUpdateOne {
	usuo.mutation.SetUpdatedAt(t)
	return usuo
}

// SetUpdatedByID sets the "updated_by_id" field.
func (usuo *UserStatusUpdateOne) SetUpdatedByID(i int) *UserStatusUpdateOne {
	usuo.mutation.SetUpdatedByID(i)
	return usuo
}

// SetNillableUpdatedByID sets the "updated_by_id" field if the given value is not nil.
func (usuo *UserStatusUpdateOne) SetNillableUpdatedByID(i *int) *UserStatusUpdateOne {
	if i != nil {
		usuo.SetUpdatedByID(*i)
	}
	return usuo
}

// ClearUpdatedByID clears the value of the "updated_by_id" field.
func (usuo *UserStatusUpdateOne) ClearUpdatedByID() *UserStatusUpdateOne {
	usuo.mutation.ClearUpdatedByID()
	return usuo
}

// SetActive sets the "active" field.
func (usuo *UserStatusUpdateOne) SetActive(b bool) *UserStatusUpdateOne {
	usuo.mutation.SetActive(b)
	return usuo
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (usuo *UserStatusUpdateOne) SetNillableActive(b *bool) *UserStatusUpdateOne {
	if b != nil {
		usuo.SetActive(*b)
	}
	return usuo
}

// SetDuration sets the "duration" field.
func (usuo *UserStatusUpdateOne) SetDuration(pg *pgtype.Tstzrange) *UserStatusUpdateOne {
	usuo.mutation.SetDuration(pg)
	return usuo
}

// ClearDuration clears the value of the "duration" field.
func (usuo *UserStatusUpdateOne) ClearDuration() *UserStatusUpdateOne {
	usuo.mutation.ClearDuration()
	return usuo
}

// SetUserID sets the "user_id" field.
func (usuo *UserStatusUpdateOne) SetUserID(i int) *UserStatusUpdateOne {
	usuo.mutation.SetUserID(i)
	return usuo
}

// SetCreatedBy sets the "created_by" edge to the User entity.
func (usuo *UserStatusUpdateOne) SetCreatedBy(u *User) *UserStatusUpdateOne {
	return usuo.SetCreatedByID(u.ID)
}

// SetUpdatedBy sets the "updated_by" edge to the User entity.
func (usuo *UserStatusUpdateOne) SetUpdatedBy(u *User) *UserStatusUpdateOne {
	return usuo.SetUpdatedByID(u.ID)
}

// SetUser sets the "user" edge to the User entity.
func (usuo *UserStatusUpdateOne) SetUser(u *User) *UserStatusUpdateOne {
	return usuo.SetUserID(u.ID)
}

// Mutation returns the UserStatusMutation object of the builder.
func (usuo *UserStatusUpdateOne) Mutation() *UserStatusMutation {
	return usuo.mutation
}

// ClearCreatedBy clears the "created_by" edge to the User entity.
func (usuo *UserStatusUpdateOne) ClearCreatedBy() *UserStatusUpdateOne {
	usuo.mutation.ClearCreatedBy()
	return usuo
}

// ClearUpdatedBy clears the "updated_by" edge to the User entity.
func (usuo *UserStatusUpdateOne) ClearUpdatedBy() *UserStatusUpdateOne {
	usuo.mutation.ClearUpdatedBy()
	return usuo
}

// ClearUser clears the "user" edge to the User entity.
func (usuo *UserStatusUpdateOne) ClearUser() *UserStatusUpdateOne {
	usuo.mutation.ClearUser()
	return usuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (usuo *UserStatusUpdateOne) Select(field string, fields ...string) *UserStatusUpdateOne {
	usuo.fields = append([]string{field}, fields...)
	return usuo
}

// Save executes the query and returns the updated UserStatus entity.
func (usuo *UserStatusUpdateOne) Save(ctx context.Context) (*UserStatus, error) {
	var (
		err  error
		node *UserStatus
	)
	if err := usuo.defaults(); err != nil {
		return nil, err
	}
	if len(usuo.hooks) == 0 {
		if err = usuo.check(); err != nil {
			return nil, err
		}
		node, err = usuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = usuo.check(); err != nil {
				return nil, err
			}
			usuo.mutation = mutation
			node, err = usuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(usuo.hooks) - 1; i >= 0; i-- {
			if usuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = usuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, usuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (usuo *UserStatusUpdateOne) SaveX(ctx context.Context) *UserStatus {
	node, err := usuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (usuo *UserStatusUpdateOne) Exec(ctx context.Context) error {
	_, err := usuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usuo *UserStatusUpdateOne) ExecX(ctx context.Context) {
	if err := usuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (usuo *UserStatusUpdateOne) defaults() error {
	if _, ok := usuo.mutation.UpdatedAt(); !ok {
		if userstatus.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized userstatus.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := userstatus.UpdateDefaultUpdatedAt()
		usuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (usuo *UserStatusUpdateOne) check() error {
	if _, ok := usuo.mutation.CreatedByID(); usuo.mutation.CreatedByCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"created_by\"")
	}
	if _, ok := usuo.mutation.UserID(); usuo.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (usuo *UserStatusUpdateOne) sqlSave(ctx context.Context) (_node *UserStatus, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userstatus.Table,
			Columns: userstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userstatus.FieldID,
			},
		},
	}
	id, ok := usuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing UserStatus.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := usuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userstatus.FieldID)
		for _, f := range fields {
			if !userstatus.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userstatus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := usuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := usuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userstatus.FieldUpdatedAt,
		})
	}
	if value, ok := usuo.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: userstatus.FieldActive,
		})
	}
	if value, ok := usuo.mutation.Duration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: userstatus.FieldDuration,
		})
	}
	if usuo.mutation.DurationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: userstatus.FieldDuration,
		})
	}
	if usuo.mutation.CreatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userstatus.CreatedByTable,
			Columns: []string{userstatus.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := usuo.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userstatus.CreatedByTable,
			Columns: []string{userstatus.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if usuo.mutation.UpdatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userstatus.UpdatedByTable,
			Columns: []string{userstatus.UpdatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := usuo.mutation.UpdatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userstatus.UpdatedByTable,
			Columns: []string{userstatus.UpdatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if usuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userstatus.UserTable,
			Columns: []string{userstatus.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := usuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userstatus.UserTable,
			Columns: []string{userstatus.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserStatus{config: usuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, usuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userstatus.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
