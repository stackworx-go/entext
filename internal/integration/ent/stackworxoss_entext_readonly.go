// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// ReadonlyClient is the client that holds all readonly ent builders.
type ReadonlyClient struct {
	client *Client
	config
	// Group is the client for interacting with the Group builders.
	Group ReadonlyGroupClient
	// Pet is the client for interacting with the Pet builders.
	Pet ReadonlyPetClient
	// User is the client for interacting with the User builders.
	User ReadonlyUserClient
	// UserStatus is the client for interacting with the UserStatus builders.
	UserStatus ReadonlyUserStatusClient
}

var ErrReadOnly = errors.New("readonly only driver")

type readonlyDriver struct{ dialect.Driver }

func (r *readonlyDriver) Exec(context.Context, string, interface{}, interface{}) error {
	return ErrReadOnly
}

func (r *readonlyDriver) BeginTx(ctx context.Context, opts *sql.TxOptions) (dialect.Tx, error) {
	return r.Driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
}

// NewReadonlyClient returns a read only client
func NewReadonlyClient(opts ...Option) *ReadonlyClient {
	client := NewClient(opts...)
	client.driver = &readonlyDriver{client.driver}
	return &ReadonlyClient{
		client:     client,
		config:     client.config,
		Group:      client.Group,
		Pet:        client.Pet,
		User:       client.User,
		UserStatus: client.UserStatus,
	}
}

func (c *ReadonlyClient) Use(hooks ...Hook) {
	c.client.Use(hooks...)
}

// GroupClient is a client for the Group schema.
type ReadonlyGroupClient interface {
	Query() *GroupQuery
	Get(ctx context.Context, id int) (*Group, error)
	GetX(ctx context.Context, id int) *Group
	QueryUsers(gr *Group) *UserQuery
}

// PetClient is a client for the Pet schema.
type ReadonlyPetClient interface {
	Query() *PetQuery
	Get(ctx context.Context, id int) (*Pet, error)
	GetX(ctx context.Context, id int) *Pet
	QueryOwner(pe *Pet) *UserQuery
}

// UserClient is a client for the User schema.
type ReadonlyUserClient interface {
	Query() *UserQuery
	Get(ctx context.Context, id int) (*User, error)
	GetX(ctx context.Context, id int) *User
	QueryStatuses(u *User) *UserStatusQuery
	QueryPets(u *User) *PetQuery
	QueryGroups(u *User) *GroupQuery
}

// UserStatusClient is a client for the UserStatus schema.
type ReadonlyUserStatusClient interface {
	Query() *UserStatusQuery
	Get(ctx context.Context, id int) (*UserStatus, error)
	GetX(ctx context.Context, id int) *UserStatus
	QueryCreatedBy(us *UserStatus) *UserQuery
	QueryUpdatedBy(us *UserStatus) *UserQuery
	QueryUser(us *UserStatus) *UserQuery
}

type ReadonlyTx struct {
	tx *Tx
	// Group is the client for interacting with the Group builders.
	Group ReadonlyGroupClient
	// Pet is the client for interacting with the Pet builders.
	Pet ReadonlyPetClient
	// User is the client for interacting with the User builders.
	User ReadonlyUserClient
	// UserStatus is the client for interacting with the UserStatus builders.
	UserStatus ReadonlyUserStatusClient
}

// Commit commits the transaction.
func (tx *ReadonlyTx) Commit() error {
	return tx.tx.Commit()
}

// OnCommit adds a hook to call on commit.
func (tx *ReadonlyTx) OnCommit(f CommitHook) {
	tx.tx.OnCommit(f)
}

func (c *ReadonlyClient) BeginTx(ctx context.Context, opts *sql.TxOptions) (*ReadonlyTx, error) {
	if opts == nil {
		opts = &sql.TxOptions{}
	}
	tx, err := c.client.BeginTx(ctx, &sql.TxOptions{
		Isolation: opts.Isolation,
		ReadOnly:  true,
	})
	if err != nil {
		return nil, err
	}
	return &ReadonlyTx{
		tx:         tx,
		Group:      tx.Group,
		Pet:        tx.Pet,
		User:       tx.User,
		UserStatus: tx.UserStatus,
	}, nil
}