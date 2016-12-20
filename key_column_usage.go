package sqlreflect

type KeyColumnUsage struct {
	ConstraintLocator
	TableLocator
	ColumnName                 string `sql:"column_name"`
	OrdinalPosition            int    `sql:"ordinal_position"`
	PositionInUniqueConstraint int    `sql:"position_in_unique_constraint"`
}
