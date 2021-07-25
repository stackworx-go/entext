package internal

import (
	"embed"

	"entgo.io/ent/entc/gen"
)

//go:embed template/*.tmpl
var templates embed.FS

var (
	EdgeHasOneTemplate = parse("template/stackworxoss_entext_edgehasone.tmpl")

	PointInTimeDeactivateTemplate = parse("template/stackworxoss_entext_pointintimedeactivate.tmpl")

	ReadonlyTemplate = parse("template/stackworxoss_entext_readonly.tmpl")

	// AllTemplates holds all templates for extending ent to support GraphQL.
	AllTemplates = []*gen.Template{
		EdgeHasOneTemplate,
		PointInTimeDeactivateTemplate,
		ReadonlyTemplate,
	}
)

func parse(name string) *gen.Template {
	bytes, err := templates.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return gen.MustParse(gen.NewTemplate(name).
		Funcs(gen.Funcs).
		Parse(string(bytes)))
}
