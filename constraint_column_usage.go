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
	ConstraintCatalog string `sql:"constraint_catalog"`
	ConstraintSchema  string `sql:"constraint_schema"`
	ConstraintName    string `sql:"constraint_name"`
}

type ConstraintColumnUsage struct {
	TableLocator
	ColumnName string `sql:"column_name"`
	ConstraintLocator
}
