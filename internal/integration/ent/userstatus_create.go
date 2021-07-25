// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/jackc/pgtype"
	"github.com/stackworx-go/entext/internal/integration/ent/user"
	"github.com/stackworx-go/entext/internal/integration/ent/userstatus"
)

// UserStatusCreate is the builder for creating a UserStatus entity.
type UserStatusCreate struct {
	config
	mutation *UserStatusMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (usc *UserStatusCreate) SetCreatedAt(t time.Time) *UserStatusCreate {
	usc.mutation.SetCreatedAt(t)
	return usc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (usc *UserStatusCreate) SetNillableCreatedAt(t *time.Time) *UserStatusCreate {
	if t != nil {
		usc.SetCreatedAt(*t)
	}
	return usc
}

// SetCreatedByID sets the "created_by_id" field.
func (usc *UserStatusCreate) SetCreatedByID(i int) *UserStatusCreate {
	usc.mutation.SetCreatedByID(i)
	return usc
}

// SetUpdatedAt sets the "updated_at" field.
func (usc *UserStatusCreate) SetUpdatedAt(t time.Time) *UserStatusCreate {
	usc.mutation.SetUpdatedAt(t)
	return usc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (usc *UserStatusCreate) SetNillableUpdatedAt(t *time.Time) *UserStatusCreate {
	if t != nil {
		usc.SetUpdatedAt(*t)
	}
	return usc
}

// SetUpdatedByID sets the "updated_by_id" field.
func (usc *UserStatusCreate) SetUpdatedByID(i int) *UserStatusCreate {
	usc.mutation.SetUpdatedByID(i)
	return usc
}

// SetNillableUpdatedByID sets the "updated_by_id" field if the given value is not nil.
func (usc *UserStatusCreate) SetNillableUpdatedByID(i *int) *UserStatusCreate {
	if i != nil {
		usc.SetUpdatedByID(*i)
	}
	return usc
}

// SetActive sets the "active" field.
func (usc *UserStatusCreate) SetActive(b bool) *UserStatusCreate {
	usc.mutation.SetActive(b)
	return usc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (usc *UserStatusCreate) SetNillableActive(b *bool) *UserStatusCreate {
	if b != nil {
		usc.SetActive(*b)
	}
	return usc
}

// SetDuration sets the "duration" field.
func (usc *UserStatusCreate) SetDuration(pg *pgtype.Tstzrange) *UserStatusCreate {
	usc.mutation.SetDuration(pg)
	return usc
}

// SetUserID sets the "user_id" field.
func (usc *UserStatusCreate) SetUserID(i int) *UserStatusCreate {
	usc.mutation.SetUserID(i)
	return usc
}

// SetCreatedBy sets the "created_by" edge to the User entity.
func (usc *UserStatusCreate) SetCreatedBy(u *User) *UserStatusCreate {
	return usc.SetCreatedByID(u.ID)
}

// SetUpdatedBy sets the "updated_by" edge to the User entity.
func (usc *UserStatusCreate) SetUpdatedBy(u *User) *UserStatusCreate {
	return usc.SetUpdatedByID(u.ID)
}

// SetUser sets the "user" edge to the User entity.
func (usc *UserStatusCreate) SetUser(u *User) *UserStatusCreate {
	return usc.SetUserID(u.ID)
}

// Mutation returns the UserStatusMutation object of the builder.
func (usc *UserStatusCreate) Mutation() *UserStatusMutation {
	return usc.mutation
}

// Save creates the UserStatus in the database.
func (usc *UserStatusCreate) Save(ctx context.Context) (*UserStatus, error) {
	var (
		err  error
		node *UserStatus
	)
	if err := usc.defaults(); err != nil {
		return nil, err
	}
	if len(usc.hooks) == 0 {
		if err = usc.check(); err != nil {
			return nil, err
		}
		node, err = usc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = usc.check(); err != nil {
				return nil, err
			}
			usc.mutation = mutation
			if node, err = usc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(usc.hooks) - 1; i >= 0; i-- {
			if usc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = usc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, usc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (usc *UserStatusCreate) SaveX(ctx context.Context) *UserStatus {
	v, err := usc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (usc *UserStatusCreate) defaults() error {
	if _, ok := usc.mutation.CreatedAt(); !ok {
		if userstatus.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized userstatus.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := userstatus.DefaultCreatedAt()
		usc.mutation.SetCreatedAt(v)
	}
	if _, ok := usc.mutation.UpdatedAt(); !ok {
		if userstatus.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized userstatus.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := userstatus.DefaultUpdatedAt()
		usc.mutation.SetUpdatedAt(v)
	}
	if _, ok := usc.mutation.Active(); !ok {
		v := userstatus.DefaultActive
		usc.mutation.SetActive(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (usc *UserStatusCreate) check() error {
	if _, ok := usc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := usc.mutation.CreatedByID(); !ok {
		return &ValidationError{Name: "created_by_id", err: errors.New(`ent: missing required field "created_by_id"`)}
	}
	if _, ok := usc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := usc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`ent: missing required field "active"`)}
	}
	if _, ok := usc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "user_id"`)}
	}
	if _, ok := usc.mutation.CreatedByID(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New("ent: missing required edge \"created_by\"")}
	}
	if _, ok := usc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New("ent: missing required edge \"user\"")}
	}
	return nil
}

func (usc *UserStatusCreate) sqlSave(ctx context.Context) (*UserStatus, error) {
	_node, _spec := usc.createSpec()
	if err := sqlgraph.CreateNode(ctx, usc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (usc *UserStatusCreate) createSpec() (*UserStatus, *sqlgraph.CreateSpec) {
	var (
		_node = &UserStatus{config: usc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: userstatus.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userstatus.FieldID,
			},
		}
	)
	if value, ok := usc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userstatus.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := usc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userstatus.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := usc.mutation.Active(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: userstatus.FieldActive,
		})
		_node.Active = value
	}
	if value, ok := usc.mutation.Duration(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: userstatus.FieldDuration,
		})
		_node.Duration = value
	}
	if nodes := usc.mutation.CreatedByIDs(); len(nodes) > 0 {
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
		_node.CreatedByID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := usc.mutation.UpdatedByIDs(); len(nodes) > 0 {
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
		_node.UpdatedByID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := usc.mutation.UserIDs(); len(nodes) > 0 {
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
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserStatusCreateBulk is the builder for creating many UserStatus entities in bulk.
type UserStatusCreateBulk struct {
	config
	builders []*UserStatusCreate
}

// Save creates the UserStatus entities in the database.
func (uscb *UserStatusCreateBulk) Save(ctx context.Context) ([]*UserStatus, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uscb.builders))
	nodes := make([]*UserStatus, len(uscb.builders))
	mutators := make([]Mutator, len(uscb.builders))
	for i := range uscb.builders {
		func(i int, root context.Context) {
			builder := uscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserStatusMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uscb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uscb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, uscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uscb *UserStatusCreateBulk) SaveX(ctx context.Context) []*UserStatus {
	v, err := uscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
