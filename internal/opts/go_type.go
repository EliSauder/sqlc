package opts

import "github.com/sqlc-dev/sqlc/internal/plugin"

// GoType describes a driver's value type and, when needed, its nullable wrapper.
// An empty NullableType means Type already represents nullable values.
// Complete means the type already includes any array dimensions.
type GoType struct {
	Type         string
	NullableType string
	Complete     bool
}

// GoTypeMapper optionally replaces an engine's built-in Go type mapping.
// GoTypeImports maps dependency package qualifiers (including the trailing dot) to
// import paths. User overrides retain precedence over both types and imports.
type GoTypeMapper interface {
	GoType(t *plugin.TypeExpr, parameter bool) GoType
	GoTypeImports() map[string]string
}
