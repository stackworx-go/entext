package entext

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/stackworx-go/entext/internal"
)

type (
	// Extension implements the entc.Extension for providing GraphQL integration.
	Extension struct {
		entc.DefaultExtension
	}
)

func (e Extension) Templates() []*gen.Template {
	return internal.AllTemplates
}
