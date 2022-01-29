package tmpl

import (
	"embed"
	"fmt"
	"io/fs"
	"strings"
	"text/template"
)

// Embed the templates directory
//go:embed templates
var Templates embed.FS

// Collection of helper functions for templates
var BaseTemplateFuncs = template.FuncMap{}

func NewTemplate(name string, helpers template.FuncMap) (t *template.Template, err error) {
	t = template.New(name).Funcs(mergeTemplateFuncMaps(BaseTemplateFuncs, helpers))
	err = loadTemplates(Templates, t)
	return t, err
}

func mergeTemplateFuncMaps(base template.FuncMap, custom template.FuncMap) template.FuncMap {
	tfm := template.FuncMap{}
	for k, f := range base {
		tfm[k] = f
	}
	for k, f := range custom {
		tfm[k] = f
	}
	return tfm
}

// LoadTemplates loads all template files into a text/template.
// The path of template is relative to the templates directory.
func loadTemplates(src embed.FS, t *template.Template) error {
	return fs.WalkDir(src, "templates", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error walking directory %s: %w", path, err)
		}
		if d.IsDir() {
			return nil
		}

		buf, err := src.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error reading file '%s': %w", path, err)
		}

		templateName := strings.TrimPrefix(path, "templates/")
		tmpl := t.New(templateName)
		_, err = tmpl.Parse(string(buf))
		if err != nil {
			return fmt.Errorf("parsing template '%s': %w", path, err)
		}
		return nil
	})
}
