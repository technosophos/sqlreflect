package sqlreflect

// TableLocator describes common fields used to locate a table.
type TableLocator struct {
	TableCatalog string `sql:"table_catalog"`
	TableSchema  string `sql:"table_schema"`
	TableName    string `sql:"table_name"`
}
