package sqlreflect

// TODO
type TableConstraint struct {
	//TableLocator
	TableCatalog   string `stbl:"table_catalog"`
	TableSchema    string `stbl:"table_schema"`
	TableNameField string `stbl:"table_name"`

	// ConstraintLocator
	ConstraintCatalog string `sql:"constraint_catalog"`
	ConstraintSchema  string `sql:"constraint_schema"`
	ConstraintName    string `sql:"constraint_name"`

	ConstraintType    ConstraintType `stbl:"constraint_type"`
	IsDeferrable      YesNo          `stbl:"is_deferrable"`
	InitiallyDeferred YesNo          `stbl:"initially_deferred"`

	opts *DBOptions
}

func (t *TableConstraint) TableName() string {
	return "information_schema.table_constraints"
}
