package sqlreflect

type ConstraintLocator struct {
	ConstraintCatalog string `sql:"constraint_catalog"`
	ConstraintSchema  string `sql:"constraint_schema"`
	ConstraintName    string `sql:"constraint_name"`
}

type ConstraintColumnUsage struct {
	TableLocator
	ColumnName string `sql:"column_name"`
	ConstraintLocator
}
