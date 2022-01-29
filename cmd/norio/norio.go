package main

import (
	"path/filepath"

	"github.com/wontaeyang/norio/pkg/codegen"
	"github.com/wontaeyang/norio/pkg/errs"
	"github.com/wontaeyang/norio/pkg/utils"
)

func main() {
	path, err := filepath.Abs("./api/petstore.yaml")
	if err != nil {
		errs.Panic("failed to load spec path", err)
	}

	spec, err := utils.LoadSpec(path)
	if err != nil {
		errs.Panic("failed to load spec", err)
	}

	opts := codegen.Options{}
	codegen.Generate(spec, "test", opts)
}
