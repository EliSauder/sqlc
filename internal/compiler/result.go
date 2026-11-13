package compiler

import (
	"github.com/sqlc-dev/sqlc/internal/opts"
	"github.com/sqlc-dev/sqlc/internal/sql/catalog"
)

type Result struct {
	Catalog      *catalog.Catalog
	Queries      []*Query
	GoTypeMapper opts.GoTypeMapper
}
