package mssql

import (
	"strings"

	"github.com/sqlc-dev/sqlc/internal/opts"
	"github.com/sqlc-dev/sqlc/internal/plugin"
)

// GoType describes the Go representations used by the go-mssqldb binding.
// Alias types remain any until the catalog resolves them to their base types.
func (p *Parser) GoType(t *plugin.TypeExpr, parameter bool) opts.GoType {
	mapped := func(typ, nullTyp string) opts.GoType {
		return opts.GoType{Type: typ, NullableType: nullTyp, Complete: true}
	}

	switch strings.ToLower(t.GetName()) {
	case "bigint":
		return mapped("int64", "sql.NullInt64")

	case "int":
		return mapped("int32", "sql.NullInt32")

	case "smallint":
		return mapped("int16", "sql.NullInt16")

	case "tinyint":
		return mapped("uint8", "sql.NullByte")

	case "bit":
		return mapped("bool", "sql.NullBool")

	case "decimal", "numeric":
		return mapped("decimal.Decimal", "decimal.NullDecimal")

	case "money", "smallmoney":
		return mapped("mssql.Money[decimal.Decimal]", "mssql.Money[decimal.NullDecimal]")

	case "float":
		return mapped("float64", "sql.NullFloat64")

	case "real":
		return mapped("float32", "sql.NullFloat64")

	case "char", "varchar", "nchar", "nvarchar", "text", "ntext", "sysname", "xml", "json":
		return mapped("string", "sql.NullString")

	case "binary", "varbinary", "image", "rowversion", "timestamp":
		return mapped("[]byte", "")

	case "date", "datetime", "datetime2", "smalldatetime", "datetimeoffset", "time":
		return mapped("time.Time", "sql.NullTime")

	case "uniqueidentifier":
		// go-mssqldb sends a NULL mssql.NullUniqueIdentifier passed by value
		// as a malformed GUID, which SQL Server rejects with error 8009; a
		// nil *mssql.UniqueIdentifier it sends as NULL.
		return mapped("mssql.UniqueIdentifier", "*mssql.UniqueIdentifier")

	case "geography", "geometry", "hierarchyid", "vector":
		return mapped("[]byte", "")

	default:
		return mapped("any", "")
	}
}

func (p *Parser) GoTypeImports() map[string]string {
	return map[string]string{
		"decimal.": "github.com/shopspring/decimal",
		"mssql.":   "github.com/microsoft/go-mssqldb",
	}
}
