package debug

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

func Print(msg string, x interface{}) {
	fmt.Printf("DEBUG: %s\n", msg)
	fmt.Printf("DEBUG: %+v\n", x)
}

func Requests(spec *openapi3.T) {
	for _, path := range spec.Paths {
		fmt.Printf("DEBUG REQUEST: %v\n", path)
	}
}

func Schemas(spec *openapi3.T) {
	for _, schema := range spec.Components.Schemas {
		fmt.Printf("DEBUG SCHEMA: %v\n", schema)
	}
}
