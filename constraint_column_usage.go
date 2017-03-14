package sqlreflect

// ConstraintType represents a type of constraint.
//
// The defined types, expressed as constants, are CHECK, PRIMARY KEY, FOREIGN KEY, and UNIQUE.
type ConstraintType string

const (
	ConstraintCheck      ConstraintType = "CHECK"
	ConstraintPrimaryKey ConstraintType = "PRIMARY KEY"
	ConstraintForeignKey ConstraintType = "FOREIGN KEY"
	ConstraintUnique     ConstraintType = "UNIQUE"
)

type ConstraintLocator struct {
	ConstraintCatalog string `stbl:"constraint_catalog"`
	ConstraintSchema  string `stbl:"constraint_schema"`
	ConstraintName    string `stbl:"constraint_name"`
}

type ConstraintColumnUsage struct {
	//TableLocator
	TableCatalog   string `stbl:"table_catalog"`
	TableSchema    string `stbl:"table_schema"`
	TableNameField string `stbl:"table_name"`

	ColumnName string `stbl:"column_name"`

	// ConstraintLocator
	ConstraintCatalog string `stbl:"constraint_catalog"`
	ConstraintSchema  string `stbl:"constraint_schema"`
	ConstraintName    string `stbl:"constraint_name"`
}
