package generate

import "strings"

// Modify the import lines to proper import paths
var imports = strings.TrimSpace(`
package errors

import (
    liberrors "todoapp/lib/errors"
)
`)
