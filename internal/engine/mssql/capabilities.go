package mssql

import "github.com/sqlc-dev/sqlc/internal/opts"

func (p *Parser) ExecutionCapabilities(sqlPackage, sqlDriver string) opts.ExecutionCapabilities {
	if sqlPackage != "" && sqlPackage != "database/sql" {
		return opts.ExecutionCapabilities{}
	}
	if sqlDriver != "" && sqlDriver != "github.com/microsoft/go-mssqldb" {
		return opts.ExecutionCapabilities{}
	}
	// The default Go binding uses go-mssqldb, whose LastInsertId always
	// errors. Other bindings may return identities through a different API.
	return opts.ExecutionCapabilities{
		LastInsertIDUnsupportedReason: "is not supported for SQL Server, whose driver reports no last insert ID; use OUTPUT INSERTED.<column> with :one",
	}
}
