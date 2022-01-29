package codegen

import (
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/wontaeyang/norio/pkg/debug"
	"github.com/wontaeyang/norio/pkg/errs"
	"github.com/wontaeyang/norio/pkg/tmpl"
)

type Options struct {
}

func Generate(doc *openapi3.T, packageName string, opts Options) {
	debug.Print("spec paths", doc.Paths)
	debug.Print("spec schema", doc.Components.Schemas)

	// Define additional template helpers
	helpers := template.FuncMap{
		"opts": func() Options { return opts },
	}

	_, err := tmpl.NewTemplate("norio", helpers)
	if err != nil {
		errs.Panic("failed to load templates", err)
	}
}
