package sqlreflect

type ConstraintType string

const (
	ConstraintCheck      ConstraintType = "CHECK"
	ConstraintForeignKey ConstraintType = "FOREIGN KEY"
	ConstraintPrimaryKey ConstraintType = "PRIMARY KEY"
	ConstraintUnique     ConstraintType = "UNIQUE"
)

// TODO
type TableConstraint struct {
	ConstraintLocator
	TableLocator
	ConstraintType    ConstraintType `sql:"constraint_type"`
	IsDeferrable      bool           `sql:"is_deferrable"`
	InitiallyDeferred bool           `sql:"initially_deferred"`
}
