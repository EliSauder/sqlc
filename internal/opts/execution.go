package opts

// ExecutionCapabilities describes restrictions of a particular driver binding.
// Empty fields retain the generator's historical behavior.
type ExecutionCapabilities struct {
	LastInsertIDUnsupportedReason string
}

// ExecutionCapabilitiesProvider is optional: engines without one do not add
// restrictions. Package and driver describe the selected execution binding,
// not the database's ability to express equivalent SQL.
type ExecutionCapabilitiesProvider interface {
	ExecutionCapabilities(sqlPackage, sqlDriver string) ExecutionCapabilities
}
