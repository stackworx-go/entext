// Code generated by entc, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/stackworx-go/entext/internal/integration/ent/schema"
	"github.com/stackworx-go/entext/internal/integration/ent/user"
	"github.com/stackworx-go/entext/internal/integration/ent/userstatus"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userMixin := schema.User{}.Mixin()
	userMixinHooks0 := userMixin[0].Hooks()
	user.Hooks[0] = userMixinHooks0[0]
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[2].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	userstatusMixin := schema.UserStatus{}.Mixin()
	userstatusMixinHooks0 := userstatusMixin[0].Hooks()
	userstatus.Hooks[0] = userstatusMixinHooks0[0]
	userstatusMixinFields0 := userstatusMixin[0].Fields()
	_ = userstatusMixinFields0
	userstatusMixinFields1 := userstatusMixin[1].Fields()
	_ = userstatusMixinFields1
	userstatusFields := schema.UserStatus{}.Fields()
	_ = userstatusFields
	// userstatusDescCreatedAt is the schema descriptor for created_at field.
	userstatusDescCreatedAt := userstatusMixinFields0[0].Descriptor()
	// userstatus.DefaultCreatedAt holds the default value on creation for the created_at field.
	userstatus.DefaultCreatedAt = userstatusDescCreatedAt.Default.(func() time.Time)
	// userstatusDescUpdatedAt is the schema descriptor for updated_at field.
	userstatusDescUpdatedAt := userstatusMixinFields0[2].Descriptor()
	// userstatus.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	userstatus.DefaultUpdatedAt = userstatusDescUpdatedAt.Default.(func() time.Time)
	// userstatus.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	userstatus.UpdateDefaultUpdatedAt = userstatusDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userstatusDescActive is the schema descriptor for active field.
	userstatusDescActive := userstatusMixinFields1[0].Descriptor()
	// userstatus.DefaultActive holds the default value on creation for the active field.
	userstatus.DefaultActive = userstatusDescActive.Default.(bool)
}

const (
	Version = "(devel)" // Version of ent codegen.
)