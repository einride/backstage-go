// Package schema provides primitives for working with catalog entity JSON schemas.
package schema

import (
	"embed"
	"io/fs"
)

//go:embed *.schema.json
var schemasFS embed.FS

// FS returns a file system with catalog entity JSON schemas.
func FS() fs.FS {
	return schemasFS
}
