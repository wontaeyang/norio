package utils

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadSpec(t *testing.T) {
	req := require.New(t)

	t.Run("random path", func(t *testing.T) {
		_, err := LoadSpec("bingbong")
		req.Error(err)
	})

	t.Run("working path", func(t *testing.T) {
		path := "../../api/petstore.yaml"
		doc, err := LoadSpec(path)
		req.NoError(err)
		req.NoError(doc.Validate(context.Background()))
	})
}
